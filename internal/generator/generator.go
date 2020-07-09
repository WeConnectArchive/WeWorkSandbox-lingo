package generator

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
)

type Settings interface {
	RootDirectory() string
	Schemas() []string
	AllowUnsupportedColumnTypes() bool
	ReplaceFieldName(name string) string
	OverrideDBTypesToPaths() map[string]PathPackageToType
}

type Parser interface {
	Tables(ctx context.Context, schema string) (<-chan string, <-chan error)
	Columns(ctx context.Context, schema, table string) (<-chan Column, <-chan error)
	ForeignKeys(ctx context.Context, schema, table string) (<-chan ForeignKey, <-chan error)
	DBTypesToPaths() map[string]PathPackageToType
}

type ForeignKey interface {
	Name() string
}

type Column interface {
	Name() string
	Table() string
	Type() *sql.ColumnType
}

func Generate(ctx context.Context, settings Settings, parser Parser) <-chan error {
	var errChan = make(chan error)

	go func() {
		defer close(errChan)

		var dbToPath = make(DBToPathType)
		for k, v := range parser.DBTypesToPaths() {
			dbToPath[k] = v
		}
		for k, v := range settings.OverrideDBTypesToPaths() {
			dbToPath[k] = v
		}

		dbTypeToPathType := func(dbType string) (PathPackageToType, error) {
			result, ok := dbToPath[dbType]
			if !ok {
				if !settings.AllowUnsupportedColumnTypes() {
					return PathPackageToType{},
						fmt.Errorf("unable to find lingo path type for DB type %s", dbType)
				}
				result = PathPackageToType{pkgCoreExpPath, "Unsupported"}
			}
			return result, nil
		}

		pipeErrors(errChan, generateSchemas(ctx, settings, parser, dbTypeToPathType))
	}()
	return errChan
}

func ensureDirectoryIsClean(directory string) error {
	if err := os.RemoveAll(directory); err != nil && !os.IsNotExist(err) {
		return err
	}
	return os.MkdirAll(directory, os.ModeDir|os.ModePerm)
}

func generateSchemas(
	ctx context.Context,
	settings Settings,
	parser Parser,
	dbTypeToPathType dbTypeToPathTypeFunc,
) <-chan error {
	var errs = make(chan error)

	go func() {
		defer close(errs)

		schemas := settings.Schemas()
		if len(schemas) == 0 {
			errs <- fmt.Errorf("no schemas selected to generate")
			return
		}

		for _, schemaName := range schemas {
			rootDir := path.Clean(settings.RootDirectory())
			schemaDir := buildSchemaDir(rootDir, schemaName)
			if schemaDir == "" {
				errs <- fmt.Errorf("root directory '%s' is not a valid path", schemaDir)
				continue
			}

			if err := ensureDirectoryIsClean(schemaDir); err != nil {
				errs <- err
				continue
			}

			tableNames, tablesErr := parser.Tables(ctx, schemaName)

			if contents, err := GenerateSchema(schemaName); err != nil {
				errs <- err
				continue
			} else if err = writeSchema(rootDir, contents, schemaName); err != nil {
				errs <- err
				continue
			}

			forEachPipeErrors(tableNames, tablesErr, errs, func(tableName interface{}) {
				retErrChan := retrieveDataAndWriteTable(
					ctx,
					settings,
					parser,
					dbTypeToPathType,
					rootDir,
					schemaName,
					tableName.(string),
				)
				pipeErrors(errs, retErrChan)
			})
		}
	}()
	return errs
}

//revive:disable-next-line - Disabling max params until we can rewrite this logic.
func retrieveDataAndWriteTable(
	ctx context.Context,
	settings Settings,
	parser Parser,
	dbTypeToPathType dbTypeToPathTypeFunc,
	schemaDir, schemaName, tableName string,
) <-chan error {
	var errs = make(chan error)

	go func() {
		defer close(errs)
		log.Printf("Generating table: %s", tableName)

		cols, colErrs := parser.Columns(ctx, schemaName, tableName)

		var tInfo = tableInfo{
			name:   tableName,
			schema: schemaName,
		}
		forEachPipeErrors(cols, colErrs, errs, func(column interface{}) {
			tInfo.columns = append(tInfo.columns, column.(Column))
		})

		columns, err := convertCols(tInfo.Columns(), settings.ReplaceFieldName, dbTypeToPathType)
		if err != nil {
			errs <- fmt.Errorf("unable to convert columns: %w", err)
			return
		}

		if contents, err := GenerateTable(tInfo, columns); err != nil {
			errs <- err
			return
		} else if err = writeTable(schemaDir, contents, schemaName, tableName); err != nil {
			errs <- err
			return
		}

		if contents, err := GeneratePackageMembers(tInfo, columns); err != nil {
			errs <- err
			return
		} else if err = writePackageMembers(schemaDir, contents, schemaName, tableName); err != nil {
			errs <- err
			return
		}
	}()
	return errs
}

type tableInfo struct {
	name        string
	schema      string
	columns     []Column
	foreignKeys []ForeignKey
}

func (t tableInfo) Name() string              { return t.name }
func (t tableInfo) Schema() string            { return t.schema }
func (t tableInfo) Columns() []Column         { return t.columns }
func (t tableInfo) ForeignKeys() []ForeignKey { return t.foreignKeys }

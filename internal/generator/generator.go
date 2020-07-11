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

		pipeErrors(errChan, generateSchemas(ctx, settings, parser))
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

			if contents, err := NewSchemaInfo(schemaName).Generate(); err != nil {
				errs <- err
				continue
			} else if err = writeSchema(rootDir, schemaName, contents); err != nil {
				errs <- err
				continue
			}

			// tablesErr is a channel and will send errors the collector below
			tableNames, tablesErr := parser.Tables(ctx, schemaName)

			dbPathTypes := updatePathTypesWithSettings(parser, settings)

			forEachPipeErrors(tableNames, tablesErr, errs, func(tableName interface{}) {
				retErrChan := retrieveDataAndWriteTable(
					ctx,
					settings,
					parser,
					dbPathTypes,
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

func updatePathTypesWithSettings(parser Parser, settings Settings) DBPathTypes {
	dbToPath := NewDBPathTypes(settings.AllowUnsupportedColumnTypes())
	dbToPath.Merge(parser.DBTypesToPaths())
	dbToPath.Merge(settings.OverrideDBTypesToPaths())
	return dbToPath
}

//revive:disable-next-line - Disabling max params until we can rewrite this logic.
func retrieveDataAndWriteTable(
	ctx context.Context,
	settings Settings,
	parser Parser,
	dbPathTypes DBPathTypes,
	schemaDir, schemaName, tableName string,
) <-chan error {
	var errs = make(chan error)

	go func() {
		defer close(errs)
		log.Printf("Generating table: %s", tableName)

		chanCols, colErrs := parser.Columns(ctx, schemaName, tableName)

		var columns = make([]Column, 0, len(chanCols))
		forEachPipeErrors(chanCols, colErrs, errs, func(column interface{}) {
			columns = append(columns, column.(Column))
		})

		cols, err := convertColumns(columns, dbPathTypes, settings.ReplaceFieldName)
		if err != nil {
			errs <- err
			return
		}

		var tInfo = TableInfo{
			Name:    tableName,
			Schema:  schemaName,
			Columns: cols,
		}
		generator, err := NewTable(tInfo)
		if err != nil {
			errs <- err
			return
		}

		contents, err := generator.GenerateTable()
		if err != nil {
			errs <- err
			return
		}
		if err = writeTable(schemaDir, schemaName, tableName, contents); err != nil {
			errs <- err
			return
		}

		if contents, err = generator.GenerateExported(); err != nil {
			errs <- err
			return
		} else if err = writePackageMembers(schemaDir, schemaName, tableName, contents); err != nil {
			errs <- err
			return
		}
	}()
	return errs
}

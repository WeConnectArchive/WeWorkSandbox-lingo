package generator

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"path"
)

type Settings interface {
	RootDirectory() string
	Schemas() []string
	ReplaceFieldName(name string) string
	OverrideDBTypesToPaths() map[string][2]string
}

type Parser interface {
	Tables(ctx context.Context, schema string) (<-chan string, <-chan error)
	Columns(ctx context.Context, schema, table string) (<-chan Column, <-chan error)
	ForeignKeys(ctx context.Context, schema, table string) (<-chan ForeignKey, <-chan error)
	DBTypesToPaths() map[string][2]string
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

	rootDir, err := ensureDirectoryIsClean(settings.RootDirectory())
	if err != nil {
		errChan <- err
		close(errChan)
		return errChan
	}

	go func() {
		defer close(errChan)

		pipeErrors(errChan, generateSchemas(ctx, settings, parser, rootDir))
	}()
	return errChan
}

func ensureDirectoryIsClean(directory string) (string, error) {
	if directory == "" {
		return "", errors.New("root directory must be a valid path")
	}
	directory = path.Clean(directory)

	if err := os.RemoveAll(directory); err != nil && !os.IsNotExist(err) {
		return "", err
	}
	if err := os.MkdirAll(directory, os.ModeDir|os.ModePerm); err != nil {
		return "", err
	}
	return directory, nil
}

func generateSchemas(ctx context.Context, settings Settings, parser Parser, rootDir string) <-chan error {
	var errors = make(chan error)

	go func() {
		defer close(errors)

		for _, schemaName := range settings.Schemas() {
			tableNames, tablesErr := parser.Tables(ctx, schemaName)

			if contents, err := GenerateSchema(schemaName); err != nil {
				errors <- err
			} else if err := writeSchema(rootDir, contents, schemaName); err != nil {
				errors <- err
			}

			forEachPipeErrors(tableNames, tablesErr, errors, func(tableName interface{}) {
				pipeErrors(errors, retrieveDataAndWriteTable(ctx, settings, parser, rootDir, schemaName, tableName.(string)))
			})
		}
	}()
	return errors
}

func retrieveDataAndWriteTable(ctx context.Context, settings Settings, parser Parser, schemaDir, schemaName, tableName string) <-chan error {
	var errors = make(chan error)

	go func() {
		defer close(errors)

		cols, colErrs := parser.Columns(ctx, schemaName, tableName)

		var tInfo = tableInfo{
			name:   tableName,
			schema: schemaName,
		}
		forEachPipeErrors(cols, colErrs, errors, func(column interface{}) {
			tInfo.columns = append(tInfo.columns, column.(Column))
		})

		var combinedTypes = combineDbTypes(parser.DBTypesToPaths(), settings.OverrideDBTypesToPaths())

		log.Printf("Generating table: %s", tInfo.Name())
		if contents, err := GenerateTable(settings, tInfo, combinedTypes); err != nil {
			errors <- err
		} else if err := writeTable(schemaDir, contents, schemaName, tableName); err != nil {
			errors <- err
		}

		if contents, err := GeneratePackageMembers(settings, tInfo, combinedTypes); err != nil {
			errors <- err
		} else if err := writePackageMembers(schemaDir, contents, schemaName, tableName); err != nil {
			errors <- err
		}
	}()
	return errors
}

func combineDbTypes(dbToPath map[string][2]string, overridePathTypes map[string][2]string) DBToPathType {
	for k, v := range overridePathTypes {
		dbToPath[k] = v
	}
	return dbToPath
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

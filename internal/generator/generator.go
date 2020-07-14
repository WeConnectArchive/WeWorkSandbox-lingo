package generator

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Settings interface {
	RootDirectory() string
	Schemas() []string
	AllowUnsupportedColumnTypes() bool
	ReplaceFieldName(name string) string
	OverrideDBTypesToPaths() map[string]PathPackageToType
}

type Parser interface {
	Tables(ctx context.Context, schema string) ([]string, error)
	Columns(ctx context.Context, schema, table string) ([]Column, error)
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

func Generate(ctx context.Context, settings Settings, parser Parser) error {
	dbPathTypes := updatePathTypesWithSettings(parser, settings)

	schemas := settings.Schemas()
	if len(schemas) == 0 {
		return fmt.Errorf("no schemas selected to generate")
	}

	for _, schemaName := range schemas {
		rootDir := settings.RootDirectory()
		if contents, err := NewSchemaInfo(schemaName).Generate(); err != nil {
			return err
		} else if err = writeSchema(rootDir, schemaName, contents); err != nil {
			return err
		}

		tableNames, tablesErr := parser.Tables(ctx, schemaName)
		if tablesErr != nil {
			return tablesErr
		}

		for _, tableName := range tableNames {
			tableErr := retrieveDataAndWriteTable(ctx, settings, parser, dbPathTypes, schemaName, tableName)
			if tableErr != nil {
				return tableErr
			}
		}
	}
	return nil
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
	schemaName, tableName string,
) error {
	log.Printf("Generating table: %s", tableName)

	columns, colErr := parser.Columns(ctx, schemaName, tableName)
	if colErr != nil {
		return colErr
	}

	cols, err := convertColumns(columns, dbPathTypes, settings.ReplaceFieldName)
	if err != nil {
		return err
	}

	var tInfo = TableInfo{
		Name:    tableName,
		Schema:  schemaName,
		Columns: cols,
	}
	generator, err := NewTable(tInfo)
	if err != nil {
		return err
	}

	contents, err := generator.GenerateTable()
	if err != nil {
		return err
	}
	if err = writeTable(settings.RootDirectory(), schemaName, tableName, contents); err != nil {
		return err
	}

	if contents, err = generator.GenerateExported(); err != nil {
		return err
	}
	return writePackageMembers(settings.RootDirectory(), schemaName, tableName, contents)
}

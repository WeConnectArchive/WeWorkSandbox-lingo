package generate

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"unicode"
	"unicode/utf8"
)

type Settings interface {
	RootDirectory() string
	Schemas() []string
	TablePrefix() string
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

	prefix, _ := utf8.DecodeRuneInString(settings.TablePrefix())
	if prefix == utf8.RuneError {
		return fmt.Errorf("first rune in '%s' is not a valid utf8 sequence", settings.TablePrefix())
	}
	if !unicode.IsGraphic(prefix) {
		return fmt.Errorf("rune hex %X is not a graphic rune", prefix)
	}

	for _, schemaName := range schemas {
		tableNames, tablesErr := parser.Tables(ctx, schemaName)
		if tablesErr != nil {
			return tablesErr
		}

		for _, tableName := range tableNames {
			tableErr := retrieveDataAndWriteTable(ctx, settings, parser, dbPathTypes, schemaName, tableName, prefix)
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
	prefix rune,
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
		Prefix:  prefix,
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
	if err = writeTable(settings.RootDirectory(), schemaName, tableName, prefix, contents); err != nil {
		return err
	}

	if contents, err = generator.GenerateExported(); err != nil {
		return err
	}
	return writePackageMembers(settings.RootDirectory(), schemaName, tableName, prefix, contents)
}

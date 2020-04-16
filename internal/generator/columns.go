package generator

import (
	"strings"
)

// DBToPathType is a map of database data/column type names to an
// array of [import package path, path type]
type DBToPathType map[string][2]string

type replaceFieldName func(string) string

type column struct {
	col      Column
	replace  replaceFieldName
	dbToPath DBToPathType
}

func (c column) DatabaseName() string {
	return c.col.Name()
}

func (c column) MemberName() string {
	return c.replace(ToNonExported(strings.ToLower(c.col.Name())))
}

func (c column) MethodName() string {
	return ToExported(strings.ToLower(c.col.Name()))
}

func (c column) PathTypeName() (string, string) {
	value, ok := c.dbToPath[c.col.Type().DatabaseTypeName()]
	if !ok || value[0] == "" || value[1] == "" {
		return "", "UnknownPathType"
	}
	return value[0], value[1]
}

func (c column) NewPathTypeName() (string, string) {
	path, name := c.PathTypeName()
	return path, "New" + name
}

func convertCols(columns []Column, replace replaceFieldName, dbToPath DBToPathType) []*column {
	var cols []*column
	for _, col := range columns {
		cols = append(cols, &column{
			col:      col,
			replace:  replace,
			dbToPath: dbToPath,
		})
	}
	return cols
}

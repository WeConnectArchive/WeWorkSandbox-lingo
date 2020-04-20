package generator

import (
	"fmt"
	"strings"
)

// DBToPathType is a map of database data/column type names to an
// array of [import package path, path type]
type DBToPathType map[string]PathPackageToType

// PathPackageToType should
type PathPackageToType = [2]string

type replaceFieldName func(string) string
type dbTypeToPathTypeFunc func(dbType string) (PathPackageToType, error)

type column struct {
	col      Column
	replace  replaceFieldName
	pathType PathPackageToType
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
	return c.pathType[0], c.pathType[1]
}

func (c column) NewPathTypeName() (string, string) {
	path, name := c.PathTypeName()
	return path, "New" + name
}

func convertCols(
	columns []Column,
	replace replaceFieldName,
	dbTypeToPathType dbTypeToPathTypeFunc,
) ([]*column, error) {
	var cols []*column
	for _, col := range columns {
		value, err := dbTypeToPathType(col.Type().DatabaseTypeName())
		if err != nil {
			return nil, fmt.Errorf("unable to find lingo path type for column named %s and type %s: %w",
				col.Name(), col.Type().DatabaseTypeName(), err)
		}

		genCol := column{
			col:      col,
			replace:  replace,
			pathType: value,
		}
		cols = append(cols, &genCol)
	}
	return cols, nil
}

package generator

import (
	"fmt"
	"strings"
)

// DBToPathType is a map of database data/column type names to an
// array of [import package path, path type]
type DBToPathType map[string]PathPackageToType

// PathPackageToType should
type PathPackageToType [2]string

func (p PathPackageToType) ShortPkg() string {
	idx := strings.LastIndexAny(p[0], "/")
	if idx == -1 && idx < len(p[0])-1 {
		return p[0]
	}
	return p[0][idx+1:]
}
func (p PathPackageToType) Pkg() string  { return p[0] }
func (p PathPackageToType) Type() string { return p[1] }

// updateFunc is used to replace a given column member name. The idea is to help prevent members that are the same name
// as keywords. For example: `type` column name would have a `type` member name on the struct. This is not allowed,
// so it is replaced to `__type` in `updateFunc`.
type updateFunc func(string) string

type column struct {
	DBName     string
	FieldName  string
	MethodName string
	PathType   PathPackageToType
}

func convertColumns(columns []Column, types DBPathTypes, updateFieldName updateFunc) ([]column, error) {
	var cols = make([]column, len(columns))
	for idx, col := range columns {
		pathType, err := types.ForDBType(col.Type().DatabaseTypeName())
		if err != nil {
			return nil, fmt.Errorf("unable to find lingo path type for column named %s and type %s: %w",
				col.Name(), col.Type().DatabaseTypeName(), err)
		}

		genCol := column{
			DBName:     col.Name(),
			FieldName:  ToNonExported(updateFieldName(col.Name())),
			MethodName: ToExported(col.Name()),
			PathType:   pathType,
		}
		cols[idx] = genCol
	}
	return cols, nil
}

package generate

import (
	"fmt"
)

func NewDBPathTypes(unsupportedAllowed bool) DBPathTypes {
	return DBPathTypes{
		types:                   make(map[string]PathPackageToType),
		allowUnsupportedColType: unsupportedAllowed,
	}
}

type DBPathTypes struct {
	types                   map[string]PathPackageToType
	allowUnsupportedColType bool
}

// Merge takes the DBPathTypes and adds them to this one.
func (d DBPathTypes) Merge(otherTypes map[string]PathPackageToType) DBPathTypes {
	for k, v := range otherTypes {
		d.types[k] = v
	}
	return d
}

// ForDBType will return the PathPackageToType for the given DB type, else error if not found. path.Unsupported will be
// used if not found instead if allowUnsupportedColType is true.
func (d DBPathTypes) ForDBType(dbType string) (PathPackageToType, error) {
	result, ok := d.types[dbType]
	if !ok {
		if !d.allowUnsupportedColType {
			return PathPackageToType{}, fmt.Errorf("unable to find lingo path type for DB type %s", dbType)
		}
		result = PathPackageToType{PkgPath, "Unsupported"}
	}
	return result, nil
}

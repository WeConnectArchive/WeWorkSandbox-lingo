package generate

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/weworksandbox/lingo/internal/generator"
)

func getSettings() *settings {
	s := &settings{}
	s.rootDirectory = viper.GetString(flagDir)
	s.schemas = viper.GetStringSlice(flagSchema)
	s.driverName = viper.GetString(flagDriver)
	s.dataSourceName = viper.GetString(flagDSN)
	s.tablePrefix = viper.GetString(flagTablePrefix)
	s.allowUnsupportedColTypes = viper.GetBool(flagUnsupportedCols)

	s.replaceNames = defaultReplaceNames
	for k, v := range viper.GetStringMapString("replace_names") {
		s.replaceNames[k] = v
	}

	s.dbToPkgTypes = make(map[string]generator.PathPackageToType)
	for k, v := range viper.GetStringMapStringSlice("db_to_pkg_type") {
		if len(v) == 2 {
			s.dbToPkgTypes[strings.ToUpper(k)] = [2]string{v[0], v[1]}
		}
	}
	return s
}

var defaultReplaceNames = map[string]string{
	"append":      "__append",
	"bool":        "__bool",
	"break":       "__break",
	"byte":        "__byte",
	"cap":         "__cap",
	"case":        "__case",
	"chan":        "__chan",
	"close":       "__close",
	"complex":     "__complex",
	"complex128":  "__complex128",
	"complex64":   "__complex64",
	"const":       "__const",
	"continue":    "__continue",
	"copy":        "__copy",
	"default":     "__default",
	"defer":       "__defer",
	"delete":      "__delete",
	"else":        "__else",
	"error":       "__error",
	"fallthrough": "__fallthrough",
	"false":       "__false",
	"float32":     "__float32",
	"float64":     "__float64",
	"for":         "__for",
	"func":        "__func",
	"go":          "__go",
	"goto":        "__goto",
	"if":          "__if",
	"imag":        "__imag",
	"import":      "__import",
	"int":         "__int",
	"int16":       "__int16",
	"int32":       "__int32",
	"int64":       "__int64",
	"int8":        "__int8",
	"interface":   "__interface",
	"iota":        "__iota",
	"len":         "__len",
	"make":        "__make",
	"map":         "__map",
	"new":         "__new",
	"nil":         "__nil",
	"package":     "__package",
	"panic":       "__panic",
	"print":       "__print",
	"println":     "__println",
	"range":       "__range",
	"real":        "__real",
	"recover":     "__recover",
	"return":      "__return",
	"rune":        "__rune",
	"select":      "__select",
	"string":      "__string",
	"struct":      "__struct",
	"switch":      "__switch",
	"true":        "__true",
	"type":        "__type",
	"uint":        "__uint",
	"uint16":      "__uint16",
	"uint32":      "__uint32",
	"uint64":      "__uint64",
	"uint8":       "__uint8",
	"uintptr":     "__uintptr",
	"var":         "__var",
}

type settings struct {
	rootDirectory            string
	schemas                  []string
	driverName               string
	dataSourceName           string
	tablePrefix              string
	allowUnsupportedColTypes bool
	replaceNames             map[string]string
	dbToPkgTypes             map[string]generator.PathPackageToType
}

func (s settings) RootDirectory() string {
	return s.rootDirectory
}

func (s settings) Schemas() []string {
	return s.schemas
}

func (s settings) DriverName() string {
	return s.driverName
}

func (s settings) DataSourceName() string {
	return s.dataSourceName
}

func (s settings) TablePrefix() string {
	return s.tablePrefix
}

func (s settings) AllowUnsupportedColumnTypes() bool {
	return s.allowUnsupportedColTypes
}

func (s settings) ReplaceFieldName(name string) string {
	newName, ok := s.replaceNames[name]
	if !ok {
		return name
	}
	return newName
}

func (s settings) OverrideDBTypesToPaths() map[string]generator.PathPackageToType {
	return s.dbToPkgTypes
}

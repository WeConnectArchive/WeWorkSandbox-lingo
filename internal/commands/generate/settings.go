package generate

import (
	"strings"

	"github.com/spf13/viper"
)

func getSettings() *settings {
	s := &settings{}
	s.rootDirectory = viper.GetString("dir")
	s.schemas = viper.GetStringSlice("schemas")
	s.driverName = viper.GetString("driver")
	s.dataSourceName = viper.GetString("dsn")
	s.allowUnsupportedColTypes = viper.GetBool("allow_unsupported_column_types")

	s.replaceNames = defaultReplaceNames
	for k, v := range viper.GetStringMapString("replace_names") {
		s.replaceNames[k] = v
	}

	s.dbToPkgTypes = make(map[string][2]string)
	for k, v := range viper.GetStringMapStringSlice("db_to_pkg_type") {
		if len(v) == 2 {
			s.dbToPkgTypes[strings.ToUpper(k)] = [2]string{v[0], v[1]}
		}
	}
	return s
}

var defaultReplaceNames = map[string]string{
	"break":       "__break",
	"default":     "__default",
	"func":        "__func",
	"interface":   "__interface",
	"select":      "__select",
	"case":        "__case",
	"defer":       "__defer",
	"go":          "__go",
	"map":         "__map",
	"struct":      "__struct",
	"chan":        "__chan",
	"else":        "__else",
	"goto":        "__goto",
	"package":     "__package",
	"switch":      "__switch",
	"const":       "__const",
	"fallthrough": "__fallthrough",
	"if":          "__if",
	"range":       "__range",
	"type":        "__type",
	"continue":    "__continue",
	"for":         "__for",
	"import":      "__import",
	"return":      "__return",
	"var":         "__var",
}

type settings struct {
	rootDirectory            string
	schemas                  []string
	driverName               string
	dataSourceName           string
	allowUnsupportedColTypes bool
	replaceNames             map[string]string
	dbToPkgTypes             map[string][2]string
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

func (s settings) OverrideDBTypesToPaths() map[string][2]string {
	return s.dbToPkgTypes
}

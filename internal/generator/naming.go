package generator

import (
	. "strings"

	"github.com/iancoleman/strcase"
)

func ToTableStruct(tableName string) string {
	return BigQ(ToExported(ToLower(tableName)))
}

func ToPackageName(s string) string {
	return ToLower(LittleQ(ToNonExported(s)))
}

func ToExported(s string) string {
	return strcase.ToCamel(s)
}
func ToNonExported(s string) string {
	return strcase.ToLowerCamel(s)
}
func LittleQ(s string) string {
	return "q" + s
}
func BigQ(s string) string {
	return "Q" + s
}

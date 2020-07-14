package generator

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

func ToTableStruct(prefix rune, tableName string) string {
	return BigPrefix(prefix, ToExported(strings.ToLower(tableName)))
}

func ToPackageName(prefix rune, s string) string {
	return strings.ToLower(LittlePrefix(prefix, ToNonExported(s)))
}

func ToExported(s string) string {
	return strcase.ToCamel(s)
}
func ToNonExported(s string) string {
	return strcase.ToLowerCamel(s)
}
func LittlePrefix(prefix rune, s string) string {
	return fmt.Sprintf("%c%s", unicode.ToLower(prefix), s)
}
func BigPrefix(prefix rune, s string) string {
	return fmt.Sprintf("%c%s", unicode.ToUpper(prefix), s)
}

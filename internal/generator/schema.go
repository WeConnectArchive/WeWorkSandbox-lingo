package generator

import (
	"fmt"
	"io"
)

type SchemaInfo struct {
	GeneratedComment string
	PackageName      string
	DBName           string
	Imports          []string
}

func NewSchemaInfo(schemaName string, prefix rune) SchemaInfo {
	return SchemaInfo{
		GeneratedComment: fmt.Sprintf(fmtSchemaHeaderComment, schemaName),
		PackageName:      ToPackageName(prefix, schemaName),
		DBName:           schemaName,
		Imports: []string{
			PkgLingo,
		},
	}
}

func (s SchemaInfo) Generate() (io.Reader, error) {
	return generateFromTemplate(schemaTemplate, s)
}

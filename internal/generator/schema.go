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

func NewSchemaInfo(schema string) SchemaInfo {
	return SchemaInfo{
		GeneratedComment: fmt.Sprintf(fmtSchemaHeaderComment, schema),
		PackageName:      ToPackageName(schema),
		DBName:           schema,
		Imports: []string{
			PkgLingo,
		},
	}
}

func (s SchemaInfo) Generate() (io.Reader, error) {
	return generateFromTemplate(schemaTemplate, s)
}

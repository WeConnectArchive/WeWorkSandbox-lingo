// Code generated by Lingo for table information_schema.SCHEMATA - DO NOT EDIT

package qschemata

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QSchemata {
	return instance
}

func CatalogName() path.StringPath {
	return instance.catalogName
}

func SchemaName() path.StringPath {
	return instance.schemaName
}

func DefaultCharacterSetName() path.StringPath {
	return instance.defaultCharacterSetName
}

func DefaultCollationName() path.StringPath {
	return instance.defaultCollationName
}

func SqlPath() path.StringPath {
	return instance.sqlPath
}

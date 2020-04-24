// Code generated by Lingo for table information_schema.SCHEMA_PRIVILEGES - DO NOT EDIT

package qschemaprivileges

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QSchemaPrivileges {
	return instance
}

func Grantee() path.String {
	return instance.grantee
}

func TableCatalog() path.String {
	return instance.tableCatalog
}

func TableSchema() path.String {
	return instance.tableSchema
}

func PrivilegeType() path.String {
	return instance.privilegeType
}

func IsGrantable() path.String {
	return instance.isGrantable
}

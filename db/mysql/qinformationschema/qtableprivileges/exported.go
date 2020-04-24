// Code generated by Lingo for table information_schema.TABLE_PRIVILEGES - DO NOT EDIT

package qtableprivileges

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QTablePrivileges {
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

func TableName() path.String {
	return instance.tableName
}

func PrivilegeType() path.String {
	return instance.privilegeType
}

func IsGrantable() path.String {
	return instance.isGrantable
}

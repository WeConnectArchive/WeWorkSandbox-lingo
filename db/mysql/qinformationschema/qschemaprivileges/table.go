// Code generated by Lingo for table information_schema.SCHEMA_PRIVILEGES - DO NOT EDIT

package qschemaprivileges

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QSchemaPrivileges {
	return newQSchemaPrivileges(alias)
}

func New() QSchemaPrivileges {
	return newQSchemaPrivileges("")
}

func newQSchemaPrivileges(alias string) QSchemaPrivileges {
	q := QSchemaPrivileges{_alias: alias}
	q.grantee = path.NewStringPath(q, "GRANTEE")
	q.tableCatalog = path.NewStringPath(q, "TABLE_CATALOG")
	q.tableSchema = path.NewStringPath(q, "TABLE_SCHEMA")
	q.privilegeType = path.NewStringPath(q, "PRIVILEGE_TYPE")
	q.isGrantable = path.NewStringPath(q, "IS_GRANTABLE")
	return q
}

type QSchemaPrivileges struct {
	_alias        string
	grantee       path.StringPath
	tableCatalog  path.StringPath
	tableSchema   path.StringPath
	privilegeType path.StringPath
	isGrantable   path.StringPath
}

// core.Table Functions

func (q QSchemaPrivileges) GetColumns() []core.Column {
	return []core.Column{
		q.grantee,
		q.tableCatalog,
		q.tableSchema,
		q.privilegeType,
		q.isGrantable,
	}
}

func (q QSchemaPrivileges) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QSchemaPrivileges) GetAlias() string {
	return q._alias
}

func (q QSchemaPrivileges) GetName() string {
	return "SCHEMA_PRIVILEGES"
}

func (q QSchemaPrivileges) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QSchemaPrivileges) Grantee() path.StringPath {
	return q.grantee
}

func (q QSchemaPrivileges) TableCatalog() path.StringPath {
	return q.tableCatalog
}

func (q QSchemaPrivileges) TableSchema() path.StringPath {
	return q.tableSchema
}

func (q QSchemaPrivileges) PrivilegeType() path.StringPath {
	return q.privilegeType
}

func (q QSchemaPrivileges) IsGrantable() path.StringPath {
	return q.isGrantable
}

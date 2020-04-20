// Code generated by Lingo for table sakila.staff_list - DO NOT EDIT

package qstafflist

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QStaffList {
	return newQStaffList(alias)
}

func New() QStaffList {
	return newQStaffList("")
}

func newQStaffList(alias string) QStaffList {
	q := QStaffList{_alias: alias}
	q.id = path.NewBoolPath(q, "ID")
	q.name = path.NewStringPath(q, "name")
	q.address = path.NewStringPath(q, "address")
	q.zipCode = path.NewStringPath(q, "zip code")
	q.phone = path.NewStringPath(q, "phone")
	q.city = path.NewStringPath(q, "city")
	q.country = path.NewStringPath(q, "country")
	q.sid = path.NewBoolPath(q, "SID")
	return q
}

type QStaffList struct {
	_alias  string
	id      path.BoolPath
	name    path.StringPath
	address path.StringPath
	zipCode path.StringPath
	phone   path.StringPath
	city    path.StringPath
	country path.StringPath
	sid     path.BoolPath
}

// core.Table Functions

func (q QStaffList) GetColumns() []core.Column {
	return []core.Column{
		q.id,
		q.name,
		q.address,
		q.zipCode,
		q.phone,
		q.city,
		q.country,
		q.sid,
	}
}

func (q QStaffList) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QStaffList) GetAlias() string {
	return q._alias
}

func (q QStaffList) GetName() string {
	return "staff_list"
}

func (q QStaffList) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QStaffList) Id() path.BoolPath {
	return q.id
}

func (q QStaffList) Name() path.StringPath {
	return q.name
}

func (q QStaffList) Address() path.StringPath {
	return q.address
}

func (q QStaffList) ZipCode() path.StringPath {
	return q.zipCode
}

func (q QStaffList) Phone() path.StringPath {
	return q.phone
}

func (q QStaffList) City() path.StringPath {
	return q.city
}

func (q QStaffList) Country() path.StringPath {
	return q.country
}

func (q QStaffList) Sid() path.BoolPath {
	return q.sid
}
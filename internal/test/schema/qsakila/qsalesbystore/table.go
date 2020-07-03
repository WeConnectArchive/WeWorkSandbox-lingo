// Code generated by Lingo for table sakila.sales_by_store - DO NOT EDIT

// +build !nolingo

package qsalesbystore

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QSalesByStore {
	return newQSalesByStore(alias)
}

func New() QSalesByStore {
	return newQSalesByStore("")
}

func newQSalesByStore(alias string) QSalesByStore {
	q := QSalesByStore{_alias: alias}
	q.store = path.NewString(q, "store")
	q.manager = path.NewString(q, "manager")
	q.totalSales = path.NewBinary(q, "total_sales")
	return q
}

type QSalesByStore struct {
	_alias     string
	store      path.String
	manager    path.String
	totalSales path.Binary
}

// core.Table Functions

func (q QSalesByStore) GetColumns() []core.Column {
	return []core.Column{
		q.store,
		q.manager,
		q.totalSales,
	}
}

func (q QSalesByStore) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QSalesByStore) GetAlias() string {
	return q._alias
}

func (q QSalesByStore) GetName() string {
	return "sales_by_store"
}

func (q QSalesByStore) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QSalesByStore) Store() path.String {
	return q.store
}

func (q QSalesByStore) Manager() path.String {
	return q.manager
}

func (q QSalesByStore) TotalSales() path.Binary {
	return q.totalSales
}

// Code generated by Lingo for table sakila.store - DO NOT EDIT

// +build !nolingo

package tstore

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TStore {
	t := New()
	if alias != "" {
		t.alias = expr.Lit(alias)
	}
	return t
}

func New() TStore {
	return TStore{}
}

type TStore struct {
	alias lingo.Expression
}

// lingo.Table Functions

func (t TStore) GetTableName() string {
	return "store"
}

func (t TStore) GetColumns() []lingo.Expression {
	return []lingo.Expression{
		t.StoreId(),
		t.ManagerStaffId(),
		t.AddressId(),
		t.LastUpdate(),
	}
}

func (t TStore) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return expr.Table(t).ToSQL(d)
}

func (t TStore) GetName() lingo.Expression {
	return expr.TableName(t)
}

func (t TStore) GetAlias() lingo.Expression {
	return t.alias
}

// Column Functions

func (t TStore) StoreId() expr.Int8 {
	return expr.Column(t, expr.Lit("store_id")).ToSQL
}

func (t TStore) ManagerStaffId() expr.Int8 {
	return expr.Column(t, expr.Lit("manager_staff_id")).ToSQL
}

func (t TStore) AddressId() expr.Int16 {
	return expr.Column(t, expr.Lit("address_id")).ToSQL
}

func (t TStore) LastUpdate() expr.Time {
	return expr.Column(t, expr.Lit("last_update")).ToSQL
}
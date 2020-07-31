// Code generated by Lingo for table sakila.customer_list - DO NOT EDIT

// +build !nolingo

package tcustomerlist

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TCustomerList {
	t := New()
	if alias != "" {
		t.alias = expr.Lit(alias)
	}
	return t
}

func New() TCustomerList {
	return TCustomerList{}
}

type TCustomerList struct {
	alias lingo.Expression
}

// lingo.Table Functions

func (t TCustomerList) GetTableName() string {
	return "customer_list"
}

func (t TCustomerList) GetColumns() []lingo.Expression {
	return []lingo.Expression{
		t.Id(),
		t.Name(),
		t.Address(),
		t.ZipCode(),
		t.Phone(),
		t.City(),
		t.Country(),
		t.Notes(),
		t.SID(),
	}
}

func (t TCustomerList) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return expr.Table(t).ToSQL(d)
}

func (t TCustomerList) GetName() lingo.Expression {
	return expr.TableName(t)
}

func (t TCustomerList) GetAlias() lingo.Expression {
	return t.alias
}

// Column Functions

func (t TCustomerList) Id() expr.Int16 {
	return expr.Column(t, expr.Lit("ID")).ToSQL
}

func (t TCustomerList) Name() expr.String {
	return expr.Column(t, expr.Lit("name")).ToSQL
}

func (t TCustomerList) Address() expr.String {
	return expr.Column(t, expr.Lit("address")).ToSQL
}

func (t TCustomerList) ZipCode() expr.String {
	return expr.Column(t, expr.Lit("zip code")).ToSQL
}

func (t TCustomerList) Phone() expr.String {
	return expr.Column(t, expr.Lit("phone")).ToSQL
}

func (t TCustomerList) City() expr.String {
	return expr.Column(t, expr.Lit("city")).ToSQL
}

func (t TCustomerList) Country() expr.String {
	return expr.Column(t, expr.Lit("country")).ToSQL
}

func (t TCustomerList) Notes() expr.String {
	return expr.Column(t, expr.Lit("notes")).ToSQL
}

func (t TCustomerList) SID() expr.Int8 {
	return expr.Column(t, expr.Lit("SID")).ToSQL
}
// Code generated by Lingo for table sakila.city - DO NOT EDIT

// +build !nolingo

package tcity

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TCity {
	t := New()
	if alias != "" {
		t.alias = expr.Lit(alias)
	}
	return t
}

func New() TCity {
	return TCity{}
}

type TCity struct {
	alias lingo.Expression
}

// lingo.Table Functions

func (t TCity) GetTableName() string {
	return "city"
}

func (t TCity) GetColumns() []lingo.Expression {
	return []lingo.Expression{
		t.CityId(),
		t.City(),
		t.CountryId(),
		t.LastUpdate(),
	}
}

func (t TCity) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return expr.Table(t).ToSQL(d)
}

func (t TCity) GetName() lingo.Expression {
	return expr.TableName(t)
}

func (t TCity) GetAlias() lingo.Expression {
	return t.alias
}

// Column Functions

func (t TCity) CityId() expr.Int16 {
	return expr.Column(t, expr.Lit("city_id")).ToSQL
}

func (t TCity) City() expr.String {
	return expr.Column(t, expr.Lit("city")).ToSQL
}

func (t TCity) CountryId() expr.Int16 {
	return expr.Column(t, expr.Lit("country_id")).ToSQL
}

func (t TCity) LastUpdate() expr.Time {
	return expr.Column(t, expr.Lit("last_update")).ToSQL
}

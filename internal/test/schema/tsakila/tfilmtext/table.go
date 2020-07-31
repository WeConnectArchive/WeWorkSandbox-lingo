// Code generated by Lingo for table sakila.film_text - DO NOT EDIT

// +build !nolingo

package tfilmtext

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TFilmText {
	t := New()
	if alias != "" {
		t.alias = expr.Lit(alias)
	}
	return t
}

func New() TFilmText {
	return TFilmText{}
}

type TFilmText struct {
	alias lingo.Expression
}

// lingo.Table Functions

func (t TFilmText) GetTableName() string {
	return "film_text"
}

func (t TFilmText) GetColumns() []lingo.Expression {
	return []lingo.Expression{
		t.FilmId(),
		t.Title(),
		t.Description(),
	}
}

func (t TFilmText) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return expr.Table(t).ToSQL(d)
}

func (t TFilmText) GetName() lingo.Expression {
	return expr.TableName(t)
}

func (t TFilmText) GetAlias() lingo.Expression {
	return t.alias
}

// Column Functions

func (t TFilmText) FilmId() expr.Int16 {
	return expr.Column(t, expr.Lit("film_id")).ToSQL
}

func (t TFilmText) Title() expr.String {
	return expr.Column(t, expr.Lit("title")).ToSQL
}

func (t TFilmText) Description() expr.String {
	return expr.Column(t, expr.Lit("description")).ToSQL
}
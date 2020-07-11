// Code generated by Lingo for table sakila.film_category - DO NOT EDIT

// +build !nolingo

package qfilmcategory

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) QFilmCategory {
	return newQFilmCategory(alias)
}

func New() QFilmCategory {
	return newQFilmCategory("")
}

func newQFilmCategory(alias string) QFilmCategory {
	q := QFilmCategory{
		_alias: alias,
	}
	q.filmId = path.NewInt16(q, "film_id")
	q.categoryId = path.NewInt8(q, "category_id")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QFilmCategory struct {
	_alias string

	filmId     path.Int16
	categoryId path.Int8
	lastUpdate path.Time
}

// lingo.Table Functions

func (q QFilmCategory) GetColumns() []lingo.Column {
	return []lingo.Column{
		q.filmId,
		q.categoryId,
		q.lastUpdate,
	}
}

func (q QFilmCategory) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QFilmCategory) GetAlias() string {
	return q._alias
}

func (q QFilmCategory) GetName() string {
	return "film_category"
}

func (q QFilmCategory) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QFilmCategory) FilmId() path.Int16 {
	return q.filmId
}

func (q QFilmCategory) CategoryId() path.Int8 {
	return q.categoryId
}

func (q QFilmCategory) LastUpdate() path.Time {
	return q.lastUpdate
}

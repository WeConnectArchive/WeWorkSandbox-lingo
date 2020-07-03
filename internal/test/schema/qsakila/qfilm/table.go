// Code generated by Lingo for table sakila.film - DO NOT EDIT

// +build !nolingo

package qfilm

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QFilm {
	return newQFilm(alias)
}

func New() QFilm {
	return newQFilm("")
}

func newQFilm(alias string) QFilm {
	q := QFilm{_alias: alias}
	q.filmId = path.NewInt16(q, "film_id")
	q.title = path.NewString(q, "title")
	q.description = path.NewString(q, "description")
	q.releaseYear = path.NewUnsupported(q, "release_year")
	q.languageId = path.NewInt8(q, "language_id")
	q.originalLanguageId = path.NewInt8(q, "original_language_id")
	q.rentalDuration = path.NewInt8(q, "rental_duration")
	q.rentalRate = path.NewBinary(q, "rental_rate")
	q.length = path.NewInt16(q, "length")
	q.replacementCost = path.NewBinary(q, "replacement_cost")
	q.rating = path.NewString(q, "rating")
	q.specialFeatures = path.NewString(q, "special_features")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QFilm struct {
	_alias             string
	filmId             path.Int16
	title              path.String
	description        path.String
	releaseYear        path.Unsupported
	languageId         path.Int8
	originalLanguageId path.Int8
	rentalDuration     path.Int8
	rentalRate         path.Binary
	length             path.Int16
	replacementCost    path.Binary
	rating             path.String
	specialFeatures    path.String
	lastUpdate         path.Time
}

// core.Table Functions

func (q QFilm) GetColumns() []core.Column {
	return []core.Column{
		q.filmId,
		q.title,
		q.description,
		q.releaseYear,
		q.languageId,
		q.originalLanguageId,
		q.rentalDuration,
		q.rentalRate,
		q.length,
		q.replacementCost,
		q.rating,
		q.specialFeatures,
		q.lastUpdate,
	}
}

func (q QFilm) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QFilm) GetAlias() string {
	return q._alias
}

func (q QFilm) GetName() string {
	return "film"
}

func (q QFilm) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QFilm) FilmId() path.Int16 {
	return q.filmId
}

func (q QFilm) Title() path.String {
	return q.title
}

func (q QFilm) Description() path.String {
	return q.description
}

func (q QFilm) ReleaseYear() path.Unsupported {
	return q.releaseYear
}

func (q QFilm) LanguageId() path.Int8 {
	return q.languageId
}

func (q QFilm) OriginalLanguageId() path.Int8 {
	return q.originalLanguageId
}

func (q QFilm) RentalDuration() path.Int8 {
	return q.rentalDuration
}

func (q QFilm) RentalRate() path.Binary {
	return q.rentalRate
}

func (q QFilm) Length() path.Int16 {
	return q.length
}

func (q QFilm) ReplacementCost() path.Binary {
	return q.replacementCost
}

func (q QFilm) Rating() path.String {
	return q.rating
}

func (q QFilm) SpecialFeatures() path.String {
	return q.specialFeatures
}

func (q QFilm) LastUpdate() path.Time {
	return q.lastUpdate
}

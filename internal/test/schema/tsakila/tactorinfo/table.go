// Code generated by Lingo for table sakila.actor_info - DO NOT EDIT

// +build !nolingo

package tactorinfo

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TActorInfo {
	return newTActorInfo(alias)
}

func New() TActorInfo {
	return newTActorInfo("")
}

func newTActorInfo(alias string) TActorInfo {
	t := TActorInfo{
		_alias: alias,
	}
	t.actorId = path.NewInt16(t, "actor_id")
	t.firstName = path.NewString(t, "first_name")
	t.lastName = path.NewString(t, "last_name")
	t.filmInfo = path.NewString(t, "film_info")
	return t
}

type TActorInfo struct {
	_alias string

	actorId   path.Int16
	firstName path.String
	lastName  path.String
	filmInfo  path.String
}

// lingo.Table Functions

func (t TActorInfo) GetColumns() []lingo.Column {
	return []lingo.Column{
		t.actorId,
		t.firstName,
		t.lastName,
		t.filmInfo,
	}
}

func (t TActorInfo) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t TActorInfo) GetAlias() string {
	return t._alias
}

func (t TActorInfo) GetName() string {
	return "actor_info"
}

func (t TActorInfo) GetParent() string {
	return "sakila"
}

// Column Functions

func (t TActorInfo) ActorId() path.Int16 {
	return t.actorId
}

func (t TActorInfo) FirstName() path.String {
	return t.firstName
}

func (t TActorInfo) LastName() path.String {
	return t.lastName
}

func (t TActorInfo) FilmInfo() path.String {
	return t.filmInfo
}
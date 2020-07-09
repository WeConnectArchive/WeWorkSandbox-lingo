// Code generated by Lingo for table sakila.actor - DO NOT EDIT

// +build !nolingo

package qactor

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QActor {
	return newQActor(alias)
}

func New() QActor {
	return newQActor("")
}

func newQActor(alias string) QActor {
	q := QActor{_alias: alias}
	q.actorId = path.NewInt16(q, "actor_id")
	q.firstName = path.NewString(q, "first_name")
	q.lastName = path.NewString(q, "last_name")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QActor struct {
	_alias     string
	actorId    path.Int16
	firstName  path.String
	lastName   path.String
	lastUpdate path.Time
}

// core.Table Functions

func (q QActor) GetColumns() []core.Column {
	return []core.Column{
		q.actorId,
		q.firstName,
		q.lastName,
		q.lastUpdate,
	}
}

func (q QActor) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QActor) GetAlias() string {
	return q._alias
}

func (q QActor) GetName() string {
	return "actor"
}

func (q QActor) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QActor) ActorId() path.Int16 {
	return q.actorId
}

func (q QActor) FirstName() path.String {
	return q.firstName
}

func (q QActor) LastName() path.String {
	return q.lastName
}

func (q QActor) LastUpdate() path.Time {
	return q.lastUpdate
}

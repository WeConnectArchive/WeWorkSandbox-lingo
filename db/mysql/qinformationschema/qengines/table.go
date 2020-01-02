// Code generated by Lingo for table information_schema.ENGINES - DO NOT EDIT

package qengines

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QEngines {
	return newQEngines(alias)
}

func New() QEngines {
	return newQEngines("")
}

func newQEngines(alias string) QEngines {
	q := QEngines{_alias: alias}
	q.engine = path.NewStringPath(q, "ENGINE")
	q.support = path.NewStringPath(q, "SUPPORT")
	q.comment = path.NewStringPath(q, "COMMENT")
	q.transactions = path.NewStringPath(q, "TRANSACTIONS")
	q.xa = path.NewStringPath(q, "XA")
	q.savepoints = path.NewStringPath(q, "SAVEPOINTS")
	return q
}

type QEngines struct {
	_alias       string
	engine       path.StringPath
	support      path.StringPath
	comment      path.StringPath
	transactions path.StringPath
	xa           path.StringPath
	savepoints   path.StringPath
}

// core.Table Functions

func (q QEngines) GetColumns() []core.Column {
	return []core.Column{
		q.engine,
		q.support,
		q.comment,
		q.transactions,
		q.xa,
		q.savepoints,
	}
}

func (q QEngines) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QEngines) GetAlias() string {
	return q._alias
}

func (q QEngines) GetName() string {
	return "ENGINES"
}

func (q QEngines) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QEngines) Engine() path.StringPath {
	return q.engine
}

func (q QEngines) Support() path.StringPath {
	return q.support
}

func (q QEngines) Comment() path.StringPath {
	return q.comment
}

func (q QEngines) Transactions() path.StringPath {
	return q.transactions
}

func (q QEngines) Xa() path.StringPath {
	return q.xa
}

func (q QEngines) Savepoints() path.StringPath {
	return q.savepoints
}

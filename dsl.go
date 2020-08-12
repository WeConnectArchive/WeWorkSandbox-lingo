package lingo

import (
	"github.com/weworksandbox/lingo/sql"
)

type Name interface {
	GetName() Expression
}

type Table interface {
	Expression
	Name
	GetTableName() string
	GetAlias() Expression
	GetColumns() []Expression
}

type Column interface {
	Expression
	Name
}

// ExpressionFunc is a wrapper similar to http.HandlerFunc
type ExpressionFunc func(d Dialect) (sql.Data, error)

func (e ExpressionFunc) ToSQL(d Dialect) (sql.Data, error) { return e(d) }

// 1. Dialect - knowing how to build each sql part `{0}.{1}`
// 2. QueryContext - know about fields, tables -> exposed to be used in 'query' package

type QueryContext interface {
}

type Expression interface {
	ToSQL(d Dialect) (sql.Data, error)
}

type Set interface {
	Expression
}

type OrderBy interface {
	Expression
}

type ComboExpression interface {
	Expression
	And(Expression) ComboExpression
	Or(Expression) ComboExpression
}

type Dialect interface {
	Name
}

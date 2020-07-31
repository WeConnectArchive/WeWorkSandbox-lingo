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

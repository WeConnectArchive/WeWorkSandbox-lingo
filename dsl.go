package lingo

import (
	"github.com/weworksandbox/lingo/sql"
)

type Name interface {
	GetName() string
}

type Alias interface {
	GetAlias() string
}

type Table interface {
	Expression
	Alias
	Name
	GetColumns() []Column
	GetParent() string
}

type Column interface {
	Expression
	Alias
	Name
	GetParent() Table
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

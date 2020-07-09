package expr

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

func Star() lingo.Expression {
	return &star{}
}

type star struct{}

func (star) ToSQL(_ lingo.Dialect) (sql.Data, error) {
	return sql.String("*"), nil
}

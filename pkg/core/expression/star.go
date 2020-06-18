package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func Star() core.Expression {
	return &star{}
}

type star struct{}

func (star) ToSQL(_ core.Dialect) (sql.Data, error) {
	return sql.String("*"), nil
}

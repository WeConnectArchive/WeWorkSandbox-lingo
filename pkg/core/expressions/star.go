package expressions

import (
	"github.com/weworksandbox/lingo/pkg/core"
)

func Star() core.Expression {
	return &star{}
}

type star struct{}

func (star) GetSQL(d core.Dialect) (core.SQL, error) {
	return core.NewSQL("*", nil), nil
}

package expressions

import (
	"github.com/weworksandbox/lingo/core"
)

func Star() core.Expression {
	return &star{}
}

type star struct{}

func (star) GetSQL(d core.Dialect, sql core.SQL) error {
	sql.AppendString("*")
	return nil
}

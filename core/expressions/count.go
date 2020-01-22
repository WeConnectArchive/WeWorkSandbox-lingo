package expressions

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/helpers"
)

func Count(countOn core.Expression) core.Expression {
	return &count{
		countOn: countOn,
	}
}

type count struct {
	countOn core.Expression
}

func (c count) GetSQL(d core.Dialect, sql core.SQL) error {
	if helpers.IsValueNilOrBlank(c.countOn) {
		return expression.ExpressionIsNil("countOn")
	}

	countOnSQL := sql.New()
	countOnErr := c.countOn.GetSQL(d, countOnSQL)
	if countOnErr != nil {
		return countOnErr
	}

	sql.AppendString("COUNT").AppendSql(countOnSQL.SurroundWithParens())
	return nil
}

package expressions

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
)

func Count(countOn core.Expression) core.Expression {
	return &count{
		countOn: countOn,
	}
}

type count struct {
	countOn core.Expression
}

func (c count) GetSQL(d core.Dialect) (core.SQL, error) {
	if helpers.IsValueNilOrBlank(c.countOn) {
		return nil, expression.ExpressionIsNil("countOn")
	}

	countOn, countOnErr := c.countOn.GetSQL(d)
	if countOnErr != nil {
		return nil, countOnErr
	}

	return core.NewSQL("COUNT", nil).AppendSQL(countOn.SurroundWithParens()), nil
}

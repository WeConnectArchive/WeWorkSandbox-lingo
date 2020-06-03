package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
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
	if check.IsValueNilOrBlank(c.countOn) {
		return nil, ExpressionIsNil("countOn")
	}

	countOn, countOnErr := c.countOn.GetSQL(d)
	if countOnErr != nil {
		return nil, countOnErr
	}

	return core.NewSQL("COUNT", nil).AppendSQL(countOn.SurroundWithParens()), nil
}

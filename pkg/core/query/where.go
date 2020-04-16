package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func BuildWhereSQL(d core.Dialect, values []core.Expression) (core.SQL, error) {
	var where = core.NewSQL("WHERE ", nil)

	switch length := len(values); {
	case length == 1:
		whereSQL, err := values[0].GetSQL(d)
		if err != nil {
			return nil, err
		}
		return where.AppendSQL(whereSQL), nil

	case length > 1:
		andOperator := expression.NewOperator(values[0], operator.And, values[1:]...)
		whereSQL, err := andOperator.GetSQL(d)
		if err != nil {
			return nil, err
		}
		return where.AppendSQL(whereSQL), nil
	}
	return core.NewSQL("", nil), nil
}

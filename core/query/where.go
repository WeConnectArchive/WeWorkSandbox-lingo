package query

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/operator"
)

func BuildWhereSQL(d core.Dialect, values []core.Expression) (core.SQL, error) {
	var where = core.NewSQL("WHERE ", nil)
	switch length := len(values); {
	case length == 1:
		if whereSql, err := values[0].GetSQL(d); err != nil {
			return nil, err
		} else {
			return where.AppendSql(whereSql), nil
		}
	case length > 1:
		andOperator := expression.NewOperator(values[0], operator.And, values[1:]...)
		if whereSql, err := andOperator.GetSQL(d); err != nil {
			return nil, err
		} else {
			return where.AppendSql(whereSql), nil
		}
	}
	return core.NewSQL("", nil), nil
}

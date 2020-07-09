package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

// BuildWhereSQL is to be used by custom queries to build a WHERE clause.
func BuildWhereSQL(d core.Dialect, values []core.Expression) (sql.Data, error) {
	var where = sql.String("WHERE ")

	switch length := len(values); {
	case length == 1:
		whereSQL, err := values[0].ToSQL(d)
		if err != nil {
			return nil, err
		}
		return where.Append(whereSQL), nil

	case length > 1:
		andOperator := operator.NewOperator(values[0], operator.And, values[1:]...)
		whereSQL, err := andOperator.ToSQL(d)
		if err != nil {
			return nil, err
		}
		return where.Append(whereSQL), nil
	}
	return sql.Empty(), nil
}

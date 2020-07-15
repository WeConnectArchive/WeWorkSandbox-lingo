package query

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

// BuildWhereSQL is to be used by custom queries to build a WHERE clause.
func BuildWhereSQL(d lingo.Dialect, values []lingo.Expression) (sql.Data, error) {
	var where = sql.String("WHERE ")

	switch length := len(values); {
	case length == 1:
		whereSQL, err := values[0].ToSQL(d)
		if err != nil {
			return nil, err
		}
		return where.Append(whereSQL), nil

	case length > 1:
		andOperator := operator.NewVariadic(values[0], operator.OpAnd, values[1:])
		whereSQL, err := andOperator.ToSQL(d)
		if err != nil {
			return nil, err
		}
		return where.Append(whereSQL), nil
	}
	return sql.Empty(), nil
}

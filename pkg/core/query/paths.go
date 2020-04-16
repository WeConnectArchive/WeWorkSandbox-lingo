package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
)

func ExpandTables(paths []core.Expression) []core.Expression {
	var expanded = make([]core.Expression, 0)
	for _, singlePath := range paths {
		if entity, ok := singlePath.(core.Table); ok {
			for _, column := range entity.GetColumns() {
				expanded = append(expanded, column)
			}
		} else {
			expanded = append(expanded, singlePath)
		}
	}
	return expanded
}

// CombinePathSQL will validate each path is not nil, and will append n+1 with a comma separating them
func CombinePathSQL(d core.Dialect, paths []core.Expression) (core.SQL, error) {
	var sql = core.NewEmptySQL()
	for _, p := range paths {
		if helpers.IsValueNilOrBlank(p) {
			return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("path entry"), sql.String())
		}

		pathSQL, err := p.GetSQL(d)
		if err != nil {
			return nil, expression.ErrorAroundSQL(err, sql.String())
		}

		if _, ok := p.(*SelectQuery); ok {
			// This case helps us add parens if the path is of a query type
			sql = sql.AppendFormat(", %s", pathSQL.SurroundWithParens().String()).AppendValues(pathSQL.Values())
		} else {
			if sql.String() == "" {
				sql = sql.AppendSQL(pathSQL)
			} else {
				sql = sql.AppendFormat(", %s", pathSQL.String()).AppendValues(pathSQL.Values())
			}
		}
	}
	return sql, nil
}

// CombineSQL will validate each path is not nil, and will append each SQL to the previous
// with a single space between each.
func CombineSQL(d core.Dialect, paths []core.Expression) (core.SQL, error) {
	var sql = core.NewEmptySQL()
	for _, p := range paths {
		if helpers.IsValueNilOrBlank(p) {
			return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("path entry"), sql.String())
		}

		pathSQL, err := p.GetSQL(d)
		if err != nil {
			return nil, expression.ErrorAroundSQL(err, sql.String())
		}

		if sql.String() == "" {
			sql = sql.AppendSQL(pathSQL)
		} else {
			sql = sql.AppendSQLWithSpace(pathSQL)
		}
	}
	return sql, nil
}

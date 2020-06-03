package query

import (
	"fmt"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
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
		if check.IsValueNilOrBlank(p) {
			return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("path entry"), sql.String())
		}

		pathSQL, err := p.GetSQL(d)
		if err != nil {
			return nil, expression.ErrorAroundSQL(err, sql.String())
		}

		if sql.String() == "" {
			sql = sql.AppendSQL(pathSQL)
		} else {
			sql = sql.AppendFormat(", %s", pathSQL.String()).AppendValues(pathSQL.Values())
		}
	}
	return sql, nil
}

// CombineSQL will validate each path is not nil, and will append each SQL to the previous
// with a single space between each.
func CombineSQL(d core.Dialect, paths []core.Expression) (core.SQL, error) {
	var sql = core.NewEmptySQL()
	for idx, p := range paths {
		if check.IsValueNilOrBlank(p) {
			entry := fmt.Sprintf("path entry[%d]", idx)
			return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil(entry), sql.String())
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

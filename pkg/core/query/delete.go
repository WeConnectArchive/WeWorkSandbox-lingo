package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/join"
)

// Delete allows deletion of an entity
func Delete(t core.Table) *DeleteQuery {
	d := DeleteQuery{}
	d.from = []core.Expression{t}
	return &d
}

type DeleteQuery struct {
	from  []core.Expression
	join  []core.Expression
	where []core.Expression
}

func (d DeleteQuery) Where(exp ...core.Expression) *DeleteQuery {
	d.where = append(d.where, exp...)
	return &d
}

// DELETE w
// FROM WorkRecord2 w
// LEFT JOIN Employee e
// ON EmployeeRun=EmployeeNo
// WHERE w.Company = '1' AND e.Date = '2013-05-06'
func (d DeleteQuery) Join(left core.Expression, jt join.Type, on core.Expression) *DeleteQuery {
	d.join = append(d.join, expression.NewJoinOn(left, jt, on))
	return &d
}

func (d DeleteQuery) GetSQL(dialect core.Dialect) (core.SQL, error) {
	var sql = core.NewSQL("DELETE FROM", nil)

	if check.IsValueNilOrEmpty(d.from) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionCannotBeEmpty("from"), sql.String())
	}
	fromSQL, err := CombinePathSQL(dialect, d.from)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	}
	sql = sql.AppendSQLWithSpace(fromSQL)

	if joinSQL, err := CombineSQL(dialect, d.join); err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	} else if joinSQL.String() != "" {
		sql = sql.AppendSQLWithSpace(joinSQL)
	}

	whereSQL, err := BuildWhereSQL(dialect, d.where)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	}
	if whereSQL.String() != "" {
		sql = sql.AppendSQLWithSpace(whereSQL)
	}

	return sql, nil
}

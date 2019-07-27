package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
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
func (d DeleteQuery) Join(left core.Expression, jt expression.JoinType, on core.Expression) *DeleteQuery {
	d.join = append(d.join, expression.NewJoinOn(left, jt, on))
	return &d
}

func (d DeleteQuery) GetSQL(dialect core.Dialect) (core.SQL, error) {
	var sql = core.NewSQL("DELETE FROM", nil)

	if helpers.IsValueNilOrEmpty(d.from) {
		return nil, expression.ErrorAroundSql(expression.ExpressionCannotBeEmpty("from"), sql.String())
	}
	if fromSql, err := CombinePathSQL(dialect, d.from); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else {
		sql = sql.AppendSqlWithSpace(fromSql)
	}

	if joinSql, err := CombineSQL(dialect, d.join); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else if joinSql.String() != "" {
		sql = sql.AppendSqlWithSpace(joinSql)
	}

	if whereSql, err := BuildWhereSQL(dialect, d.where); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else if whereSql.String() != "" {
		sql = sql.AppendSqlWithSpace(whereSql)
	}

	return sql, nil
}

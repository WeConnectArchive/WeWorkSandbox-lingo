package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/join"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

// Delete allows deletion of an entity
func Delete(t core.Table) *DeleteQuery {
	return &DeleteQuery{
		from: t,
	}
}

type DeleteQuery struct {
	from  core.Expression
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

func (d DeleteQuery) ToSQL(dialect core.Dialect) (sql.Data, error) {
	var s = sql.String("DELETE FROM")

	if check.IsValueNilOrBlank(d.from) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("from"), s.String())
	}

	from, err := d.from.ToSQL(dialect)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	s = s.AppendWithSpace(from)

	if joinSQL, err := JoinToSQL(dialect, sepSpace, d.join); err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	} else if joinSQL.String() != "" {
		s = s.AppendWithSpace(joinSQL)
	}

	whereSQL, err := BuildWhereSQL(dialect, d.where)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	if whereSQL.String() != "" {
		s = s.AppendWithSpace(whereSQL)
	}

	return s, nil
}

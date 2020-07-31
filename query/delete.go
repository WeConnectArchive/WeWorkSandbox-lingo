package query

import (
	"errors"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/join"
	"github.com/weworksandbox/lingo/sql"
)

// Delete allows deletion of an entity
func Delete(from lingo.Table) *DeleteQuery {
	return &DeleteQuery{
		from: from.GetName(),
	}
}

type DeleteQuery struct {
	from  lingo.Expression
	join  []lingo.Expression
	where lingo.Expression
}

func (d DeleteQuery) Where(exp ...lingo.ComboExpression) *DeleteQuery {
	d.where = appendCombosWith(d.where, exp, expr.And)
	return &d
}

// DELETE w
// FROM WorkRecord2 w
// LEFT JOIN Employee e
// ON EmployeeRun=EmployeeNo
// WHERE w.Company = '1' AND e.Date = '2013-05-06'
func (d DeleteQuery) Join(left lingo.Expression, jt join.Type, on lingo.Expression) *DeleteQuery {
	d.join = append(d.join, join.NewJoinOn(left, jt, on))
	return &d
}

func (d DeleteQuery) ToSQL(dialect lingo.Dialect) (sql.Data, error) {
	var s = sql.String("DELETE FROM")

	if check.IsValueNilOrBlank(d.from) {
		return nil, NewErrAroundSQL(s, errors.New("from cannot be empty"))
	}

	from, err := d.from.ToSQL(dialect)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	s = s.AppendWithSpace(from)

	if joinSQL, err := JoinToSQL(dialect, sepSpace, d.join); err != nil {
		return nil, NewErrAroundSQL(s, err)
	} else if joinSQL.String() != "" {
		s = s.AppendWithSpace(joinSQL)
	}

	whereSQL, err := buildIfNotEmpty(dialect, d.where)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	if whereSQL.String() != "" {
		s = s.AppendWithSpace(sqlWhere).AppendWithSpace(whereSQL)
	}

	return s, nil
}

package query

import (
	"errors"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func Update(table core.Table) *UpdateQuery {
	update := UpdateQuery{
		table: table,
	}
	return &update
}

type UpdateQuery struct {
	table core.Table
	set   []core.Expression
	where []core.Expression
}

func (u UpdateQuery) Where(exp ...core.Expression) *UpdateQuery {
	u.where = append(u.where, exp...)
	return &u
}

func (u UpdateQuery) Set(exp ...core.Set) *UpdateQuery {
	u.set = append(u.set, castSetsToExpressions(exp)...)
	return &u
}

func (u UpdateQuery) ToSQL(d core.Dialect) (sql.Data, error) {
	var s = sql.String("UPDATE")

	if check.IsValueNilOrBlank(u.table) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("table"), s.String())
	}
	if u.table.GetAlias() != "" {
		return nil, expression.ErrorAroundSQL(errors.New("table alias must be unset"), s.String())
	}
	table, err := u.table.ToSQL(d)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	s = s.AppendWithSpace(table)

	if check.IsValueNilOrEmpty(u.set) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("set"), s.String())
	}
	pathsSQL, err := JoinToSQL(d, sepPathComma, u.set)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	if pathsSQL.String() != "" {
		s = s.AppendWithSpace(sql.String("SET")).AppendWithSpace(pathsSQL)
	}

	whereSQL, err := BuildWhereSQL(d, u.where)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	if whereSQL.String() != "" {
		s = s.AppendWithSpace(whereSQL)
	}

	return s, nil
}

func castSetsToExpressions(sets []core.Set) []core.Expression {
	var exp = make([]core.Expression, 0, len(sets))
	for _, set := range sets {
		exp = append(exp, set)
	}
	return exp
}

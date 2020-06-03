package query

import (
	"errors"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/check"
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

func (u UpdateQuery) GetSQL(d core.Dialect) (core.SQL, error) {
	var sql = core.NewSQL("UPDATE", nil)

	if check.IsValueNilOrBlank(u.table) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("table"), sql.String())
	}
	if u.table.GetAlias() != "" {
		return nil, expression.ErrorAroundSQL(errors.New("table alias must be unset"), sql.String())
	}
	tableSQL, err := u.table.GetSQL(d)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	}
	sql = sql.AppendSQLWithSpace(tableSQL)

	if check.IsValueNilOrEmpty(u.set) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("set"), sql.String())
	}
	pathsSQL, err := CombinePathSQL(d, u.set)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	}
	if pathsSQL.String() != "" {
		sql = sql.AppendStringWithSpace("SET").AppendSQLWithSpace(pathsSQL)
	}

	whereSQL, err := BuildWhereSQL(d, u.where)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	}
	if whereSQL.String() != "" {
		sql = sql.AppendSQLWithSpace(whereSQL)
	}

	return sql, nil
}

func castSetsToExpressions(sets []core.Set) []core.Expression {
	var exp = make([]core.Expression, 0, len(sets))
	for _, set := range sets {
		exp = append(exp, set)
	}
	return exp
}

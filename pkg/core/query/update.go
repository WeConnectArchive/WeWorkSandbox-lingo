package query

import (
	"errors"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
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

	if helpers.IsValueNilOrBlank(u.table) {
		return nil, expression.ErrorAroundSql(expression.ExpressionIsNil("table"), sql.String())
	}
	if u.table.GetAlias() != "" {
		return nil, expression.ErrorAroundSql(errors.New("table alias must be unset"), sql.String())
	}
	if tableSql, err := u.table.GetSQL(d); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else {
		sql = sql.AppendSqlWithSpace(tableSql)
	}

	if helpers.IsValueNilOrEmpty(u.set) {
		return nil, expression.ErrorAroundSql(expression.ExpressionIsNil("set"), sql.String())
	}
	if pathsSql, err := CombinePathSQL(d, u.set); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else if pathsSql.String() != "" {
		sql = sql.AppendStringWithSpace("SET").AppendSqlWithSpace(pathsSql)
	}

	if whereSql, err := BuildWhereSQL(d, u.where); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else if whereSql.String() != "" {
		sql = sql.AppendSqlWithSpace(whereSql)
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

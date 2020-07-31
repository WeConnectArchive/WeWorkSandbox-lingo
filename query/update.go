package query

import (
	"errors"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func Update(table lingo.Table) *UpdateQuery {
	update := UpdateQuery{
		table: table.GetName(),
	}
	return &update
}

type UpdateQuery struct {
	table lingo.Expression
	set   []lingo.Expression
	where lingo.Expression
}

func (u UpdateQuery) Where(exp ...lingo.ComboExpression) *UpdateQuery {
	u.where = appendCombosWith(u.where, exp, expr.And)
	return &u
}

func (u UpdateQuery) Set(exp ...lingo.Set) *UpdateQuery {
	u.set = append(u.set, castSetsToExpressions(exp)...)
	return &u
}

func (u UpdateQuery) ToSQL(d lingo.Dialect) (sql.Data, error) {
	var s = sql.String("UPDATE")

	if check.IsValueNilOrBlank(u.table) {
		return nil, NewErrAroundSQL(s, errors.New("table cannot be empty"))
	}
	table, err := u.table.ToSQL(d)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	s = s.AppendWithSpace(table)

	if check.IsValueNilOrEmpty(u.set) {
		return nil, NewErrAroundSQL(s, errors.New("set cannot be empty"))
	}
	pathsSQL, err := JoinToSQL(d, sepPathComma, u.set)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	if pathsSQL.String() != "" {
		s = s.AppendWithSpace(sql.String("SET")).AppendWithSpace(pathsSQL)
	}

	whereSQL, err := buildIfNotEmpty(d, u.where)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	if whereSQL.String() != "" {
		s = s.AppendWithSpace(sqlWhere).AppendWithSpace(whereSQL)
	}

	return s, nil
}

func castSetsToExpressions(sets []lingo.Set) []lingo.Expression {
	var exp = make([]lingo.Expression, 0, len(sets))
	for _, set := range sets {
		exp = append(exp, set)
	}
	return exp
}

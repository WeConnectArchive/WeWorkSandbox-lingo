package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

func Select(paths ...core.Expression) *SelectQuery {
	selectPart := SelectQuery{}
	selectPart.paths = paths
	return &selectPart
}

func SelectFrom(e core.Table) *SelectQuery {
	return Select(e).From(e)
}

type SelectQuery struct {
	from  core.Expression
	join  []core.Expression
	where []core.Expression
	order []core.Expression
	paths []core.Expression
}

func (s *SelectQuery) From(e core.Table) *SelectQuery {
	s.from = e
	return s
}

func (s *SelectQuery) Where(exp ...core.Expression) *SelectQuery {
	s.where = append(s.where, exp...)
	return s
}

func (s *SelectQuery) OrderBy(exp core.Expression, direction sort.Direction) *SelectQuery {
	s.order = append(s.order, expression.NewOrderBy(exp, direction))
	return s
}

func (s *SelectQuery) Join(left core.Expression, jt expression.JoinType, on core.Expression) *SelectQuery {
	s.join = append(s.join, expression.NewJoinOn(left, jt, on))
	return s
}

func (s *SelectQuery) GetSQL(d core.Dialect) (core.SQL, error) {
	var sql = core.NewSQL("SELECT ", nil)

	if helpers.IsValueNilOrEmpty(s.paths) {
		return nil, expression.ErrorAroundSql(expression.ExpressionCannotBeEmpty("columns"), sql.String())
	}
	if pathsSql, err := CombinePathSQL(d, ExpandTables(s.paths)); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else {
		sql = sql.AppendSqlWithSpace(pathsSql)
	}

	if helpers.IsValueNilOrBlank(s.from) {
		return nil, expression.ExpressionIsNil("from")
	}
	if fromSql, err := s.from.GetSQL(d); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else {
		sql = sql.AppendStringWithSpace("FROM").AppendSqlWithSpace(fromSql)
	}

	if joinSql, err := CombineSQL(d, s.join); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else if joinSql.String() != "" {
		sql = sql.AppendSqlWithSpace(joinSql)
	}

	if whereSql, err := BuildWhereSQL(d, s.where); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else if whereSql.String() != "" {
		sql = sql.AppendSqlWithSpace(whereSql)
	}

	if orderBySql, err := CombinePathSQL(d, s.order); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else if orderBySql.String() != "" {
		sql = sql.AppendStringWithSpace("ORDER BY").AppendSqlWithSpace(orderBySql)
	}

	return sql, nil
}

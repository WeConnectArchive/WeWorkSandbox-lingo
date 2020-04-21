package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
	"github.com/weworksandbox/lingo/pkg/core/join"
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

func (s *SelectQuery) Join(left core.Expression, jt join.Type, on core.Expression) *SelectQuery {
	s.join = append(s.join, expression.NewJoinOn(left, jt, on))
	return s
}

func (s *SelectQuery) GetSQL(d core.Dialect) (core.SQL, error) {
	sql, err := s.selectFrom(d)
	if err != nil {
		return nil, err
	}

	fromSQL, err := s.from.GetSQL(d)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	}
	sql = sql.AppendStringWithSpace("FROM").AppendSQLWithSpace(fromSQL)

	if joinSQL, err := CombineSQL(d, s.join); err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	} else if joinSQL.String() != "" {
		sql = sql.AppendSQLWithSpace(joinSQL)
	}

	if whereSQL, err := BuildWhereSQL(d, s.where); err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	} else if whereSQL.String() != "" {
		sql = sql.AppendSQLWithSpace(whereSQL)
	}

	if orderBySQL, err := CombinePathSQL(d, s.order); err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	} else if orderBySQL.String() != "" {
		sql = sql.AppendStringWithSpace("ORDER BY").AppendSQLWithSpace(orderBySQL)
	}

	return sql, nil
}

func (s *SelectQuery) selectFrom(d core.Dialect) (core.SQL, error) {
	var sql = core.NewSQL("SELECT ", nil)
	if helpers.IsValueNilOrEmpty(s.paths) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionCannotBeEmpty("columns"), sql.String())
	}
	pathsSQL, err := CombinePathSQL(d, ExpandTables(s.paths))
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, sql.String())
	}
	sql = sql.AppendSQLWithSpace(pathsSQL)

	if helpers.IsValueNilOrBlank(s.from) {
		return nil, expression.ExpressionIsNil("from")
	}
	return sql, nil
}

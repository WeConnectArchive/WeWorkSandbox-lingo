package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/join"
	"github.com/weworksandbox/lingo/pkg/core/expression/sort"
	"github.com/weworksandbox/lingo/pkg/core/sql"
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
	from     core.Expression
	join     []core.Expression
	where    []core.Expression
	order    []core.Expression
	paths    []core.Expression
	modifier Modifier
}

func (q *SelectQuery) From(e core.Table) *SelectQuery {
	q.from = e
	return q
}

func (q *SelectQuery) Where(exp ...core.Expression) *SelectQuery {
	q.where = append(q.where, exp...)
	return q
}

func (q *SelectQuery) OrderBy(exp core.Expression, direction sort.Direction) *SelectQuery {
	q.order = append(q.order, sort.NewOrderBy(exp, direction))
	return q
}

// Join an expression with a specific joinType using an on statement.
func (q *SelectQuery) Join(left core.Expression, joinType join.Type, on core.Expression) *SelectQuery {
	q.join = append(q.join, join.NewJoinOn(left, joinType, on))
	return q
}

// Restrict the query with things like limits and offsets.
func (q *SelectQuery) Restrict(m Modifier) *SelectQuery {
	q.modifier = m
	return q
}

func (q *SelectQuery) ToSQL(d core.Dialect) (sql.Data, error) {
	s, err := q.selectFrom(d)
	if err != nil {
		return nil, err
	}

	from, err := q.from.ToSQL(d)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	s = s.AppendWithSpace(sql.String("FROM")).AppendWithSpace(from)

	if joinSQL, err := JoinToSQL(d, sepSpace, q.join); err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	} else if joinSQL.String() != "" {
		s = s.AppendWithSpace(joinSQL)
	}

	if where, err := BuildWhereSQL(d, q.where); err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	} else if where.String() != "" {
		s = s.AppendWithSpace(where)
	}

	if orderBy, err := JoinToSQL(d, sepPathComma, q.order); err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	} else if orderBy.String() != "" {
		s = s.AppendWithSpace(sql.String("ORDER BY")).AppendWithSpace(orderBy)
	}

	if s, err = q.buildModifier(d, s); err != nil {
		return nil, err // Already wrapped
	}
	return s, nil
}

func (q *SelectQuery) selectFrom(d core.Dialect) (sql.Data, error) {
	var s = sql.String("SELECT")
	if check.IsValueNilOrEmpty(q.paths) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionCannotBeEmpty("columns"), s.String())
	}
	pathsSQL, err := JoinToSQL(d, sepPathComma, ExpandTables(q.paths))
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	s = s.AppendWithSpace(pathsSQL)

	if check.IsValueNilOrBlank(q.from) {
		return nil, expression.ExpressionIsNil("from")
	}
	return s, nil
}

// buildModifier will determine if the modifier was set / needs to be built, and return the resulting SQL. This will
// check if the dialect support Modify on queries, if it is not & a modifier was set, it errors.
func (q *SelectQuery) buildModifier(d core.Dialect, s sql.Data) (sql.Data, error) {
	if q.modifier == nil || q.modifier.IsZero() {
		return s, nil
	}

	modifyDialect, ok := d.(Modify)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("Modify")
	}

	modify, err := modifyDialect.Modify(q.modifier)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	return s.AppendWithSpace(modify), nil
}

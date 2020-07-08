package query

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
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
		return nil, err // Already wrapped
	}

	if check.IsValueNilOrBlank(q.from) {
		return nil, NewErrAroundSQL(s, errors.New("from cannot be empty"))
	}
	from, err := q.from.ToSQL(d)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	s = s.AppendWithSpace(sql.String("FROM")).AppendWithSpace(from)

	if joinSQL, err := JoinToSQL(d, sepSpace, q.join); err != nil {
		return nil, NewErrAroundSQL(s, err)
	} else if joinSQL.String() != "" {
		s = s.AppendWithSpace(joinSQL)
	}

	if where, err := BuildWhereSQL(d, q.where); err != nil {
		return nil, NewErrAroundSQL(s, err)
	} else if where.String() != "" {
		s = s.AppendWithSpace(where)
	}

	if orderBy, err := JoinToSQL(d, sepPathComma, q.order); err != nil {
		return nil, NewErrAroundSQL(s, err)
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
		return nil, NewErrAroundSQL(s, errors.New("columns cannot be empty"))
	}
	pathsSQL, err := JoinToSQL(d, sepPathComma, ExpandTables(q.paths))
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	return s.AppendWithSpace(pathsSQL), nil
}

// buildModifier will determine if the modifier was set / needs to be built, and return the resulting SQL. This will
// check if the dialect support ModifyDialect on queries, if it is not & a modifier was set, it errors.
func (q *SelectQuery) buildModifier(d core.Dialect, s sql.Data) (sql.Data, error) {
	if q.modifier == nil || q.modifier.IsZero() {
		return s, nil
	}

	modifyDialect, ok := d.(ModifyDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'ModifyDialect'", d.GetName())
	}

	modify, err := modifyDialect.Modify(q.modifier)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	return s.AppendWithSpace(modify), nil
}

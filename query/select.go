package query

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/join"
	"github.com/weworksandbox/lingo/query/sort"
	"github.com/weworksandbox/lingo/sql"
)

func Select(paths ...lingo.Expression) *SelectQuery {
	selectPart := SelectQuery{}
	selectPart.paths = paths
	return &selectPart
}

func SelectFrom(e lingo.Table) *SelectQuery {
	return Select(e.GetColumns()...).From(e.GetName())
}

type SelectQuery struct {
	from     lingo.Expression
	join     []lingo.Expression
	where    lingo.Expression
	order    []lingo.Expression
	paths    []lingo.Expression
	modifier Modifier
}

func (q *SelectQuery) From(e lingo.Expression) *SelectQuery {
	q.from = e
	return q
}

func (q *SelectQuery) Where(exp ...lingo.ComboExpression) *SelectQuery {
	q.where = appendCombosWith(q.where, exp, expr.And)
	return q
}

func (q *SelectQuery) OrderBy(exp lingo.Expression, direction sort.Direction) *SelectQuery {
	q.order = append(q.order, sort.NewOrderBy(exp, direction))
	return q
}

// Join an expr with a specific joinType using an on statement.
func (q *SelectQuery) Join(left lingo.Expression, joinType join.Type, on lingo.Expression) *SelectQuery {
	q.join = append(q.join, join.NewJoinOn(left, joinType, on))
	return q
}

// Restrict the query with things like limits and offsets.
func (q *SelectQuery) Restrict(m Modifier) *SelectQuery {
	q.modifier = m
	return q
}

func (q *SelectQuery) ToSQL(d lingo.Dialect) (sql.Data, error) {
	s, err := q.selectFrom(d)
	if err != nil {
		return nil, err // Already wrapped
	}

	s, err = q.buildFrom(d, s)
	if err != nil {
		return nil, err // Already wrapped
	}

	if joinSQL, err := JoinToSQL(d, sepSpace, q.join); err != nil {
		return nil, NewErrAroundSQL(s, err)
	} else if joinSQL.String() != "" {
		s = s.AppendWithSpace(joinSQL)
	}

	if where, err := buildIfNotEmpty(d, q.where); err != nil {
		return nil, NewErrAroundSQL(s, err)
	} else if where.String() != "" {
		s = s.AppendWithSpace(sqlWhere).AppendWithSpace(where)
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

func (q *SelectQuery) selectFrom(d lingo.Dialect) (sql.Data, error) {
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

func (q *SelectQuery) buildFrom(d lingo.Dialect, s sql.Data) (sql.Data, error) {
	if check.IsValueNilOrBlank(q.from) {
		return nil, NewErrAroundSQL(s, errors.New("from cannot be empty"))
	}
	from, err := q.from.ToSQL(d)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	s = s.AppendWithSpace(sql.String("FROM")).AppendWithSpace(from)
	return s, nil
}

// buildModifier will determine if the modifier was set / needs to be built, and return the resulting SQL. This will
// check if the dialect support ModifyDialect on queries, if it is not & a modifier was set, it errors.
func (q *SelectQuery) buildModifier(d lingo.Dialect, s sql.Data) (sql.Data, error) {
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

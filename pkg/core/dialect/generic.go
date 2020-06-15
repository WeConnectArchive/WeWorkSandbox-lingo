package dialect

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/join"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

var genericJoinTypeToStr = map[join.Type]string{
	join.Inner: "INNER JOIN",
	join.Outer: "OUTER JOIN",
	join.Left:  "LEFT JOIN",
	join.Right: "RIGHT JOIN",
}

// AliasElseName will use the core.Alias if non-empty, else the Name is used.
func AliasElseName(n core.Name) core.SQL {
	alias, ok := n.(core.Alias)
	if aliasStr := alias.GetAlias(); ok && aliasStr != "" {
		return core.NewSQLString(aliasStr)
	}
	return core.NewSQLString(n.GetName())
}

func ExpandTable(entity core.Table) (core.SQL, error) {
	sql := core.NewSQLString(entity.GetName())
	if alias := entity.GetAlias(); alias != "" {
		sql = sql.AppendFormat(" AS %s", alias)
	}
	return sql, nil
}

func ExpandTableWithSchema(entity core.Table) (core.SQL, error) {
	sql, err := ExpandTable(entity)
	if err != nil {
		return nil, fmt.Errorf("unable to expand table before schema: %w", err)
	}
	return core.NewSQLf("%s.", entity.GetParent()).AppendSQL(sql), nil
}

func ExpandColumn(column core.Column) (core.SQL, error) {
	sql := core.NewSQLString(column.GetName())
	if a := column.GetAlias(); a != "" {
		sql = sql.AppendFormat(" AS %s", a)
	}
	return sql, nil
}

func ExpandColumnWithParent(column core.Column) (core.SQL, error) {
	table := AliasElseName(column.GetParent())
	colSQL, err := ExpandColumn(column)
	if err != nil {
		return nil, fmt.Errorf("unable to expand column: %w", err)
	}
	// Append separator prior to column: `table.column`
	return table.AppendString(".").AppendSQL(colSQL), nil
}

type ValueFormatter interface {
	ValueFormat(count int) core.SQL
}

func Value(formatter ValueFormatter, values []interface{}) (core.SQL, error) {
	if check.IsValueNilOrBlank(values) {
		return nil, expression.ConstantIsNil()
	}
	if check.IsValueNilOrBlank(formatter) {
		return nil, errors.New("ValueFormatter is nil or the interface pointer is nil")
	}

	return formatter.ValueFormat(len(values)).AppendValues(values), nil
}

func Operator(left core.SQL, op operator.Operand, values []core.SQL) (core.SQL, error) {
	opWithSpaces := " " + op.String() + " "

	switch op {
	case operator.And, operator.Or:
		return left.CombineWithSeparator(values, opWithSpaces).SurroundWithParens(), nil
	case operator.Eq, operator.NotEq, operator.LessThan, operator.LessThanOrEqual,
		operator.GreaterThan, operator.GreaterThanOrEqual, operator.Like, operator.NotLike,
		operator.Between, operator.NotBetween:
		return left.CombineWithSeparator(values, opWithSpaces), nil
	case operator.In, operator.NotIn:
		sql := core.NewEmptySQL().CombinePaths(values).SurroundWithParens()
		return left.AppendStringWithSpace(op.String()).AppendSQLWithSpace(sql), nil
	case operator.Null, operator.NotNull:
		return left.AppendStringWithSpace(op.String()), nil
	}
	return nil, expression.ErrorAroundSQL(expression.EnumIsInvalid("Operator", op), left.String())
}

func Join(left core.SQL, joinType string, on core.SQL) (core.SQL, error) {
	if check.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	if check.IsValueNilOrBlank(on.String()) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("on"), left.String())
	}
	if check.IsValueNilOrEmpty(joinType) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("joinType"), left.String())
	}

	var sql = core.NewSQL(joinType, nil)
	return sql.AppendSQLWithSpace(left).AppendString(" ON").AppendSQLWithSpace(on), nil
}

type SetFormatter interface {
	SetValueFormat() string
}

func Set(format SetFormatter, left core.SQL, value core.SQL) (core.SQL, error) {
	if check.IsValueNilOrBlank(format) {
		return nil, errors.New("SetFormatter is nil or the interface pointer is nil")
	}
	if check.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	if check.IsValueNilOrBlank(value.String()) {
		return nil, expression.ExpressionIsNil("value")
	}

	return left.AppendStringWithSpace(format.SetValueFormat()).AppendSQLWithSpace(value), nil
}

func OrderBy(left core.SQL, direction sort.Direction) (core.SQL, error) {
	if check.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	switch direction {
	case sort.Ascending, sort.Descending:
		return left.AppendStringWithSpace(direction.String()), nil
	}
	return nil, expression.EnumIsInvalid("direction", direction)
}

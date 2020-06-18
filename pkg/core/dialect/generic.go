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
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var genericJoinTypeToStr = map[join.Type]string{
	join.Inner: "INNER JOIN",
	join.Outer: "OUTER JOIN",
	join.Left:  "LEFT JOIN",
	join.Right: "RIGHT JOIN",
}

// AliasElseName will use the core.Alias if non-empty, else the Name is used.
func AliasElseName(n core.Name) sql.Data {
	alias, ok := n.(core.Alias)
	if aliasStr := alias.GetAlias(); ok && aliasStr != "" {
		return sql.String(aliasStr)
	}
	return sql.String(n.GetName())
}

func ExpandTable(entity core.Table) (sql.Data, error) {
	s := sql.String(entity.GetName())
	if alias := entity.GetAlias(); alias != "" {
		s = s.Append(sql.Format(" AS %s", alias))
	}
	return s, nil
}

func ExpandTableWithSchema(entity core.Table) (sql.Data, error) {
	s, err := ExpandTable(entity)
	if err != nil {
		return nil, fmt.Errorf("unable to expand table before schema: %w", err)
	}
	return sql.Format("%s.", entity.GetParent()).Append(s), nil
}

func ExpandColumn(column core.Column) (sql.Data, error) {
	s := sql.String(column.GetName())
	if a := column.GetAlias(); a != "" {
		s = s.Append(sql.Format(" AS %s", a))
	}
	return s, nil
}

func ExpandColumnWithParent(column core.Column) (sql.Data, error) {
	table := AliasElseName(column.GetParent())
	colSQL, err := ExpandColumn(column)
	if err != nil {
		return nil, fmt.Errorf("unable to expand column: %w", err)
	}
	// Append separator prior to column: `table.column`
	return table.Append(sql.String(".").Append(colSQL)), nil
}

type ValueFormatter interface {
	ValueFormat(count int) sql.Data
}

func Value(formatter ValueFormatter, values []interface{}) (sql.Data, error) {
	if check.IsValueNilOrBlank(values) {
		return nil, expression.ConstantIsNil()
	}
	if check.IsValueNilOrBlank(formatter) {
		return nil, errors.New("ValueFormatter is nil or the interface pointer is nil")
	}

	return formatter.ValueFormat(len(values)).Append(sql.Values(values)), nil
}

func Operator(left sql.Data, op operator.Operand, values []sql.Data) (sql.Data, error) {
	opWithSpaces := " " + op.String() + " "

	switch op {

	case operator.And, operator.Or:
		// Create the where SQL and then put parens around it.
		whereSQL := sql.Join(opWithSpaces, append([]sql.Data{left}, values...))
		return sql.Surround("(", ")", whereSQL), nil

	case operator.Eq, operator.NotEq, operator.LessThan, operator.LessThanOrEqual,
		operator.GreaterThan, operator.GreaterThanOrEqual, operator.Like, operator.NotLike,
		operator.Between, operator.NotBetween:

		combined := append([]sql.Data{left}, values...)
		return sql.Join(opWithSpaces, combined), nil

	case operator.In, operator.NotIn:
		opSQL := sql.String(op.String())
		return left.AppendWithSpace(opSQL).SurroundAppend(" (", ")", sql.Join(", ", values)), nil

	case operator.Null, operator.NotNull:
		opSQL := sql.String(op.String())
		return left.AppendWithSpace(opSQL), nil
	}
	return nil, expression.ErrorAroundSQL(expression.EnumIsInvalid("Operator", op), left.String())
}

func Join(left sql.Data, joinType string, on sql.Data) (sql.Data, error) {
	if check.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	if check.IsValueNilOrBlank(on.String()) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("on"), left.String())
	}
	if check.IsValueNilOrEmpty(joinType) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("joinType"), left.String())
	}

	return sql.String(joinType).
		AppendWithSpace(left).
		AppendWithSpace(sql.String("ON")).
		AppendWithSpace(on), nil
}

type SetFormatter interface {
	SetValueFormat() string
}

func Set(format SetFormatter, left sql.Data, value sql.Data) (sql.Data, error) {
	if check.IsValueNilOrBlank(format) {
		return nil, errors.New("SetFormatter is nil or the interface pointer is nil")
	}
	if check.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	if check.IsValueNilOrBlank(value.String()) {
		return nil, expression.ExpressionIsNil("value")
	}

	return left.AppendWithSpace(sql.String(format.SetValueFormat())).AppendWithSpace(value), nil
}

func OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error) {
	if check.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	switch direction {
	case sort.Ascending, sort.Descending:
		return left.AppendWithSpace(sql.String(direction.String())), nil
	}
	return nil, expression.EnumIsInvalid("direction", direction)
}

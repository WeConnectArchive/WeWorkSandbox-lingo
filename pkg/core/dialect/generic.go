package dialect

import (
	"github.com/weworksandbox/lingo/pkg/core/helpers"
	"github.com/weworksandbox/lingo/pkg/core/sort"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"golang.org/x/xerrors"
)

var genericJoinTypeToStr = map[expression.JoinType]string{
	expression.InnerJoin: "INNER JOIN",
	expression.OuterJoin: "OUTER JOIN",
	expression.LeftJoin:  "LEFT JOIN",
	expression.RightJoin: "RIGHT JOIN",
}

func ExpandEntity(entity core.Table) (core.SQL, error) {
	sql := core.NewSQLf("%s.%s", entity.GetParent(), entity.GetName())
	if entity.GetAlias() != "" {
		return sql.AppendFormat(" AS %s", entity.GetAlias()), nil
	}
	return sql, nil
}

func ExpandColumn(column core.Column) (core.SQL, error) {
	if column.GetAlias() != "" {
		return core.NewSQLf("%s.%s AS %s", column.GetParent().GetAlias(), column.GetName(), column.GetAlias()), nil
	}
	if column.GetParent().GetAlias() != "" {
		return core.NewSQLf("%s.%s", column.GetParent().GetAlias(), column.GetName()), nil
	}
	return core.NewSQLf("%s.%s", column.GetParent().GetName(), column.GetName()), nil
}

type ValueFormatter interface {
	ValueFormat() string
}

func Value(format ValueFormatter, value interface{}) (core.SQL, error) {
	if helpers.IsValueNilOrBlank(value) {
		return nil, expression.ConstantIsNil()
	}
	if helpers.IsValueNilOrBlank(format) {
		return nil, xerrors.New("ValueFormatter is nil or the interface pointer is nil")
	}
	return core.NewSQL(format.ValueFormat(), []interface{}{value}), nil
}

func Operator(left core.SQL, op operator.Operand, values []core.SQL) (core.SQL, error) {
	switch op {
	case operator.And, operator.Or:
		return left.CombineWithSeparator(values, " "+op.String()+" ").SurroundWithParens(), nil
	case operator.Eq, operator.NotEq, operator.LessThan, operator.LessThanOrEqual,
		operator.GreaterThan, operator.GreaterThanOrEqual, operator.Like, operator.NotLike,
		operator.Between, operator.NotBetween:
		return left.CombineWithSeparator(values, " "+op.String()+" "), nil
	case operator.In, operator.NotIn:
		sql := core.NewEmptySql().CombinePaths(values).SurroundWithParens()
		return left.AppendStringWithSpace(op.String()).AppendSqlWithSpace(sql), nil
	case operator.Null, operator.NotNull:
		return left.AppendStringWithSpace(op.String()), nil
	}
	return nil, expression.ErrorAroundSql(expression.EnumIsInvalid("Operator", op), left.String())
}

func Join(left core.SQL, joinType string, on core.SQL) (core.SQL, error) {
	if helpers.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	if helpers.IsValueNilOrBlank(on.String()) {
		return nil, expression.ErrorAroundSql(expression.ExpressionIsNil("on"), left.String())
	}
	if helpers.IsValueNilOrEmpty(joinType) {
		return nil, expression.ErrorAroundSql(expression.ExpressionIsNil("joinType"), left.String())
	}

	var sql = core.NewSQL(joinType, nil)
	return sql.AppendSqlWithSpace(left).AppendString(" ON").AppendSqlWithSpace(on), nil
}

type SetFormatter interface {
	SetValueFormat() string
}

func Set(format SetFormatter, left core.SQL, value core.SQL) (core.SQL, error) {
	if helpers.IsValueNilOrBlank(format) {
		return nil, xerrors.New("SetFormatter is nil or the interface pointer is nil")
	}
	if helpers.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	if helpers.IsValueNilOrBlank(value.String()) {
		return nil, expression.ExpressionIsNil("value")
	}

	return left.AppendStringWithSpace(format.SetValueFormat()).AppendSqlWithSpace(value), nil
}

func OrderBy(left core.SQL, direction sort.Direction) (core.SQL, error) {
	if helpers.IsValueNilOrBlank(left.String()) {
		return nil, expression.ExpressionIsNil("left")
	}
	switch direction {
	case sort.Ascending, sort.Descending:
		return left.AppendStringWithSpace(direction.String()), nil
	}
	return nil, expression.EnumIsInvalid("direction", direction)
}

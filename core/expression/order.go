package expression

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/sort"
)

type Order interface {
	OrderBy(left core.SQL, direction sort.Direction) error
}

func NewOrderBy(left core.Expression, direction sort.Direction) core.OrderBy {
	e := &orderBy{
		left:      left,
		direction: direction,
	}
	e.exp = e
	return e
}

type orderBy struct {
	ComboExpression
	left      core.Expression
	direction sort.Direction
}

func (o orderBy) GetSQL(d core.Dialect, sql core.SQL) error {
	order, ok := d.(Order)
	if !ok {
		return DialectFunctionNotSupported("Order")
	}

	if o.left == nil {
		return ExpressionIsNil("left")
	}

	if lerr := o.left.GetSQL(d, sql); lerr != nil {
		return lerr
	}

	if o.direction == sort.Unknown {
		return ErrorAroundSql(ExpressionIsNil("direction"), sql.String())
	}

	return order.OrderBy(sql, o.direction)
}

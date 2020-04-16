package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

type Order interface {
	OrderBy(left core.SQL, direction sort.Direction) (core.SQL, error)
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

func (o orderBy) GetSQL(d core.Dialect) (core.SQL, error) {
	order, ok := d.(Order)
	if !ok {
		return nil, DialectFunctionNotSupported("Order")
	}

	if o.left == nil {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := o.left.GetSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	if o.direction == sort.Unknown {
		return nil, ErrorAroundSQL(ExpressionIsNil("direction"), left.String())
	}

	return order.OrderBy(left, o.direction)
}

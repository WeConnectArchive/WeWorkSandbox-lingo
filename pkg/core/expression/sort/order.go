package sort

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Order interface {
	OrderBy(left sql.Data, direction Direction) (sql.Data, error)
}

func NewOrderBy(left core.Expression, direction Direction) core.OrderBy {
	e := &orderBy{
		left:      left,
		direction: direction,
	}
	e.ComboExpression = expression.NewComboExpression(e)
	return e
}

type orderBy struct {
	expression.ComboExpression
	left      core.Expression
	direction Direction
}

func (o orderBy) ToSQL(d core.Dialect) (sql.Data, error) {
	order, ok := d.(Order)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("Order")
	}

	if o.left == nil {
		return nil, expression.ExpressionIsNil("left")
	}
	left, lerr := o.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	if o.direction == Unknown {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("direction"), left.String())
	}

	return order.OrderBy(left, o.direction)
}

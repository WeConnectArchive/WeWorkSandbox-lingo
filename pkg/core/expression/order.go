package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression/sort"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Order interface {
	OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error)
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

func (o orderBy) ToSQL(d core.Dialect) (sql.Data, error) {
	order, ok := d.(Order)
	if !ok {
		return nil, DialectFunctionNotSupported("Order")
	}

	if o.left == nil {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := o.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	if o.direction == sort.Unknown {
		return nil, ErrorAroundSQL(ExpressionIsNil("direction"), left.String())
	}

	return order.OrderBy(left, o.direction)
}

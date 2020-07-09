package sort

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Dialect interface {
	OrderBy(left sql.Data, direction Direction) (sql.Data, error)
}

func NewOrderBy(left core.Expression, direction Direction) core.OrderBy {
	return orderBy{
		left:      left,
		direction: direction,
	}
}

type orderBy struct {
	left      core.Expression
	direction Direction
}

func (o orderBy) ToSQL(d core.Dialect) (sql.Data, error) {
	order, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'sort.Dialect'", d.GetName())
	}

	if o.left == nil {
		return nil, errors.New("left of 'order by' cannot be empty")
	}
	left, lerr := o.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}
	return order.OrderBy(left, o.direction)
}

package sort

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

type Dialect interface {
	OrderBy(left sql.Data, direction Direction) (sql.Data, error)
}

func NewOrderBy(left lingo.Expression, direction Direction) By {
	return By{
		left:      left,
		direction: direction,
	}
}

type By struct {
	left      lingo.Expression
	direction Direction
}

func (o By) ToSQL(d lingo.Dialect) (sql.Data, error) {
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

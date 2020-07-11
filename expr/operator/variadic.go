package operator

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/sql"
)

func NewVariadic(left lingo.Expression, op Operator, expressions []lingo.Expression) Variadic {
	return Variadic{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
}

type Variadic struct {
	left        lingo.Expression
	operand     Operator
	expressions []lingo.Expression
}

func (o Variadic) And(exp lingo.Expression) lingo.ComboExpression {
	return NewBinary(o, And, exp)
}

func (o Variadic) Or(exp lingo.Expression) lingo.ComboExpression {
	return NewBinary(o, Or, exp)
}

func (o Variadic) ToSQL(d lingo.Dialect) (sql.Data, error) {
	operand, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'operator.Dialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(o.left) {
		return nil, errors.New("left of operator.Variadic cannot be empty")
	}
	left, lerr := o.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	var sqlArr = make([]sql.Data, 0, len(o.expressions))
	for index, ex := range o.expressions {
		if check.IsValueNilOrEmpty(ex) {
			return nil, fmt.Errorf("expressions[%d] of operator.Variadic cannot be empty", index)
		}

		s, err := ex.ToSQL(d)
		if err != nil {
			return nil, err
		}
		sqlArr = append(sqlArr, s)
	}

	return operand.VariadicOperator(left, o.operand, sqlArr)
}

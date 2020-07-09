package operator

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/sql"
)

type Dialect interface {
	Operator(left sql.Data, op Operand, values []sql.Data) (sql.Data, error)
}

func NewOperator(left lingo.Expression, op Operand, expressions ...lingo.Expression) lingo.ComboExpression {
	return operate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
}

type operate struct {
	left        lingo.Expression
	operand     Operand
	expressions []lingo.Expression
}

func (o operate) And(exp lingo.Expression) lingo.ComboExpression {
	return NewOperator(o, And, exp)
}

func (o operate) Or(exp lingo.Expression) lingo.ComboExpression {
	return NewOperator(o, Or, exp)
}

func (o operate) ToSQL(d lingo.Dialect) (sql.Data, error) {
	operand, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'json.Dialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(o.left) {
		return nil, errors.New("left of operator cannot be empty")
	}
	left, lerr := o.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	var sqlArr = make([]sql.Data, 0, len(o.expressions))
	for index, ex := range o.expressions {
		if check.IsValueNilOrEmpty(ex) {
			return nil, fmt.Errorf("expressions[%d] of operator cannot be empty", index)
		}

		s, err := ex.ToSQL(d)
		if err != nil {
			return nil, err
		}
		sqlArr = append(sqlArr, s)
	}

	return operand.Operator(left, o.operand, sqlArr)
}

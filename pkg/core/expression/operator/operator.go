package operator

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type OperatorDialect interface {
	Operator(left sql.Data, op Operand, values []sql.Data) (sql.Data, error)
}

func NewOperator(left core.Expression, op Operand, expressions ...core.Expression) core.ComboExpression {
	return operate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
}

type operate struct {
	left        core.Expression
	operand     Operand
	expressions []core.Expression
}

func (o operate) And(exp core.Expression) core.ComboExpression {
	return NewOperator(o, And, exp)
}

func (o operate) Or(exp core.Expression) core.ComboExpression {
	return NewOperator(o, Or, exp)
}

func (o operate) ToSQL(d core.Dialect) (sql.Data, error) {
	operand, ok := d.(OperatorDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'OperatorDialect'", d.GetName())
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

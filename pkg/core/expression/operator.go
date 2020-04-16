package expression

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

type Operator interface {
	Operator(left core.SQL, op operator.Operand, values []core.SQL) (core.SQL, error)
}

func NewOperator(left core.Expression, op operator.Operand, expressions ...core.Expression) core.ComboExpression {
	e := &operate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
	e.exp = e
	return e
}

type operate struct {
	ComboExpression
	left        core.Expression
	operand     operator.Operand
	expressions []core.Expression
}

func (o operate) GetSQL(d core.Dialect) (core.SQL, error) {
	operand, ok := d.(Operator)
	if !ok {
		return nil, DialectFunctionNotSupported("Operator")
	}

	if helpers.IsValueNilOrEmpty(o.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := o.left.GetSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	var sqlArr = make([]core.SQL, 0, len(o.expressions))
	for index, ex := range o.expressions {
		if helpers.IsValueNilOrEmpty(ex) {
			return nil, ErrorAroundSQL(ExpressionIsNil(fmt.Sprintf("expressions[%d]", index)), left.String())
		}

		sql, err := ex.GetSQL(d)
		if err != nil {
			return nil, ErrorAroundSQL(err, left.String())
		}
		sqlArr = append(sqlArr, sql)
	}

	return operand.Operator(left, o.operand, sqlArr)
}

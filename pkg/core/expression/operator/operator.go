package operator

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Operator interface {
	Operator(left sql.Data, op Operand, values []sql.Data) (sql.Data, error)
}

func NewOperator(left core.Expression, op Operand, expressions ...core.Expression) core.ComboExpression {
	e := &operate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
	e.exp = e
	return e
}

type operate struct {
	expression.ComboExpression
	left        core.Expression
	operand     Operand
	expressions []core.Expression
}

func (o operate) ToSQL(d core.Dialect) (sql.Data, error) {
	operand, ok := d.(Operator)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("Operator")
	}

	if check.IsValueNilOrEmpty(o.left) {
		return nil, expression.ExpressionIsNil("left")
	}
	left, lerr := o.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	var sqlArr = make([]sql.Data, 0, len(o.expressions))
	for index, ex := range o.expressions {
		if check.IsValueNilOrEmpty(ex) {
			return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil(fmt.Sprintf("expressions[%d]", index)), left.String())
		}

		s, err := ex.ToSQL(d)
		if err != nil {
			return nil, expression.ErrorAroundSQL(err, left.String())
		}
		sqlArr = append(sqlArr, s)
	}

	return operand.Operator(left, o.operand, sqlArr)
}

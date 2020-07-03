package json

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Operation interface {
	JSONOperator(left sql.Data, op Operand, values []sql.Data) (sql.Data, error)
}

func NewJSONOperation(left core.Expression, op Operand, expressions ...core.Expression) core.ComboExpression {
	e := &jsonOperate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
	e.ComboExpression = expression.NewComboExpression(e)
	return e
}

type jsonOperate struct {
	expression.ComboExpression
	left        core.Expression
	operand     Operand
	expressions []core.Expression
}

func (o jsonOperate) ToSQL(d core.Dialect) (sql.Data, error) {
	operate, ok := d.(Operation)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("JSONOperation")
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

	return operate.JSONOperator(left, o.operand, sqlArr)
}

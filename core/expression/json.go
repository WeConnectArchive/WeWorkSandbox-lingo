package expression

import (
	"fmt"

	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/helpers"
	"github.com/weworksandbox/lingo/core/json"
)

type JSONOperation interface {
	JSONOperator(left core.SQL, op json.Operand, values []core.SQL) (core.SQL, error)
}

func NewJSONOperation(left core.Expression, op json.Operand, expressions ...core.Expression) core.ComboExpression {
	e := &jsonOperate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
	e.exp = e
	return e
}

type jsonOperate struct {
	ComboExpression
	left        core.Expression
	operand     json.Operand
	expressions []core.Expression
}

func (o jsonOperate) GetSQL(d core.Dialect) (core.SQL, error) {
	operate, ok := d.(JSONOperation)
	if !ok {
		return nil, DialectFunctionNotSupported("JSONOperation")
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
			return nil, ErrorAroundSql(ExpressionIsNil(fmt.Sprintf("expressions[%d]", index)), left.String())
		}

		sql, err := ex.GetSQL(d)
		if err != nil {
			return nil, ErrorAroundSql(err, left.String())
		}
		sqlArr = append(sqlArr, sql)
	}

	return operate.JSONOperator(left, o.operand, sqlArr)
}

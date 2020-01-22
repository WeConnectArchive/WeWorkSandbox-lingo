package expression

import (
	"fmt"

	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/helpers"
	"github.com/weworksandbox/lingo/core/json"
)

type JSONOperation interface {
	JSONOperator(left core.SQL, op json.Operand, values []core.SQL) error
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

func (o jsonOperate) GetSQL(d core.Dialect, sql core.SQL) error {
	operate, ok := d.(JSONOperation)
	if !ok {
		return DialectFunctionNotSupported("JSONOperation")
	}

	if helpers.IsValueNilOrEmpty(o.left) {
		return ExpressionIsNil("left")
	}
	if lerr := o.left.GetSQL(d, sql); lerr != nil {
		return lerr
	}

	var sqlArr = make([]core.SQL, 0, len(o.expressions))
	for index, ex := range o.expressions {
		if helpers.IsValueNilOrEmpty(ex) {
			return ErrorAroundSql(ExpressionIsNil(fmt.Sprintf("expressions[%d]", index)), sql.String())
		}

		expressionSQL := sql.New()
		if err := ex.GetSQL(d, expressionSQL); err != nil {
			return ErrorAroundSql(err, sql.String())
		}
		sqlArr = append(sqlArr, expressionSQL)
	}

	return operate.JSONOperator(sql, o.operand, sqlArr)
}

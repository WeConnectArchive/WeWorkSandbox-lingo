package expression

import (
	"fmt"

	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/helpers"
	"github.com/weworksandbox/lingo/core/operator"
)

type Operator interface {
	Operator(left core.SQL, op operator.Operand, values []core.SQL) error
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

func (o operate) GetSQL(d core.Dialect, sql core.SQL) error {
	operate, ok := d.(Operator)
	if !ok {
		return DialectFunctionNotSupported("Operator")
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

		expresionSQL := sql.New()
		if err := ex.GetSQL(d, expresionSQL); err != nil {
			return ErrorAroundSql(err, sql.String())
		}
		sqlArr = append(sqlArr, sql)
	}

	return operate.Operator(sql, o.operand, sqlArr)
}

package expression

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression/json"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type JSONOperation interface {
	JSONOperator(left sql.Data, op json.Operand, values []sql.Data) (sql.Data, error)
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

func (o jsonOperate) ToSQL(d core.Dialect) (sql.Data, error) {
	operate, ok := d.(JSONOperation)
	if !ok {
		return nil, DialectFunctionNotSupported("JSONOperation")
	}

	if check.IsValueNilOrEmpty(o.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := o.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	var sqlArr = make([]sql.Data, 0, len(o.expressions))
	for index, ex := range o.expressions {
		if check.IsValueNilOrEmpty(ex) {
			return nil, ErrorAroundSQL(ExpressionIsNil(fmt.Sprintf("expressions[%d]", index)), left.String())
		}

		s, err := ex.ToSQL(d)
		if err != nil {
			return nil, ErrorAroundSQL(err, left.String())
		}
		sqlArr = append(sqlArr, s)
	}

	return operate.JSONOperator(left, o.operand, sqlArr)
}

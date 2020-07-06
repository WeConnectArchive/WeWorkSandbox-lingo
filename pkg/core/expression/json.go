package expression

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/json"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type JSONOperation interface {
	JSONOperator(left sql.Data, op json.Operand, values []sql.Data) (sql.Data, error)
}

func NewJSONOperation(left core.Expression, op json.Operand, expressions ...core.Expression) core.ComboExpression {
	return jsonOperate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
}

type jsonOperate struct {
	left        core.Expression
	operand     json.Operand
	expressions []core.Expression
}

func (j jsonOperate) And(exp core.Expression) core.ComboExpression {
	return NewOperator(j, operator.And, exp)
}

func (j jsonOperate) Or(exp core.Expression) core.ComboExpression {
	return NewOperator(j, operator.Or, exp)
}

func (j jsonOperate) ToSQL(d core.Dialect) (sql.Data, error) {
	operate, ok := d.(JSONOperation)
	if !ok {
		return nil, DialectFunctionNotSupported("JSONOperation")
	}

	if check.IsValueNilOrEmpty(j.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := j.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	var sqlArr = make([]sql.Data, 0, len(j.expressions))
	for index, ex := range j.expressions {
		if check.IsValueNilOrEmpty(ex) {
			return nil, ErrorAroundSQL(ExpressionIsNil(fmt.Sprintf("expressions[%d]", index)), left.String())
		}

		s, err := ex.ToSQL(d)
		if err != nil {
			return nil, ErrorAroundSQL(err, left.String())
		}
		sqlArr = append(sqlArr, s)
	}

	return operate.JSONOperator(left, j.operand, sqlArr)
}

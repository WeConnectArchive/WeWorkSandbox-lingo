package json

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Dialect interface {
	JSONOperator(left sql.Data, op Operand, values []sql.Data) (sql.Data, error)
}

func NewJSONOperation(left core.Expression, op Operand, expressions ...core.Expression) core.ComboExpression {
	return jsonOperate{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
}

type jsonOperate struct {
	left        core.Expression
	operand     Operand
	expressions []core.Expression
}

func (j jsonOperate) And(exp core.Expression) core.ComboExpression {
	return operator.NewOperator(j, operator.And, exp)
}

func (j jsonOperate) Or(exp core.Expression) core.ComboExpression {
	return operator.NewOperator(j, operator.Or, exp)
}

func (j jsonOperate) ToSQL(d core.Dialect) (sql.Data, error) {
	operate, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'json.Dialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(j.left) {
		return nil, errors.New("left of 'json' cannot be empty")
	}
	left, lerr := j.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	var sqlArr = make([]sql.Data, 0, len(j.expressions))
	for index, ex := range j.expressions {
		if check.IsValueNilOrEmpty(ex) {
			return nil, fmt.Errorf("expressions[%d] of json cannot be empty", index)
		}

		s, err := ex.ToSQL(d)
		if err != nil {
			return nil, err
		}
		sqlArr = append(sqlArr, s)
	}

	return operate.JSONOperator(left, j.operand, sqlArr)
}

package json

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

type Dialect interface {
	JSONOperator(left sql.Data, op Operand, values []sql.Data) (sql.Data, error)
}

func NewJSONOperation(left lingo.Expression, op Operand, expressions ...lingo.Expression) Operator {
	return Operator{
		left:        left,
		operand:     op,
		expressions: expressions,
	}
}

type Operator struct {
	left        lingo.Expression
	operand     Operand
	expressions []lingo.Expression
}

func (j Operator) And(exp lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(j, operator.And, exp)
}

func (j Operator) Or(exp lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(j, operator.Or, exp)
}

func (j Operator) ToSQL(d lingo.Dialect) (sql.Data, error) {
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

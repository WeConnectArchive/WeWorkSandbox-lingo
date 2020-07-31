package expr

import (
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

type OperandDialect interface {
	BuildOperator(op Operator, operands ...lingo.Expression) (sql.Data, error)
}

// NewOperation takes an operation type and the expressions for it. The index of each expressions denotes the positional
// index of the Operator format.
func Operation(op Operator, operands ...lingo.Expression) lingo.ExpressionFunc {
	return func(d lingo.Dialect) (sql.Data, error) {
		opDialect, ok := d.(OperandDialect)
		if !ok {
			return nil, fmt.Errorf("dialect '%s' does not support 'expr.OperandDialect'", d.GetName())
		}
		return opDialect.BuildOperator(op, operands...)
	}
}

type ComboOperation lingo.ExpressionFunc

func (c ComboOperation) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return c(d)
}

func (c ComboOperation) And(exp lingo.Expression) lingo.ComboExpression {
	return And(c, exp)
}

func (c ComboOperation) Or(exp lingo.Expression) lingo.ComboExpression {
	return Or(c, exp)
}

package expr

import (
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

type OperandDialect interface {
	BuildOperator(op Operation) (sql.Data, error)
}

// NewOperation takes an operation type and the expressions for it. The index of each expressions denotes the positional
// index of the Operator format.
func NewOperation(op Operator, operands ...lingo.Expression) Operation {
	return Operation{
		Op:    op,
		Exprs: operands,
	}
}

type Operation struct {
	Op    Operator
	Exprs []lingo.Expression
}

func (o Operation) And(exp lingo.Expression) lingo.ComboExpression {
	return And(o, exp)
}

func (o Operation) Or(exp lingo.Expression) lingo.ComboExpression {
	return Or(o, exp)
}

func (o Operation) ToSQL(d lingo.Dialect) (sql.Data, error) {
	opDialect, ok := d.(OperandDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'expr.OperandDialect'", d.GetName())
	}
	return opDialect.BuildOperator(o)
}

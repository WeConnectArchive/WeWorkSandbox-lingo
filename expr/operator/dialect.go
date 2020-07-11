package operator

import (
	"github.com/weworksandbox/lingo/sql"
)

type Dialect interface {
	UnaryOperator(left sql.Data, op Operator) (sql.Data, error)
	BinaryOperator(left sql.Data, op Operator, right sql.Data) (sql.Data, error)
	VariadicOperator(left sql.Data, op Operator, values []sql.Data) (sql.Data, error)
}

package expr

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

// Lit creates a literal as a lingo.expression
func Lit(s string) lingo.Expression {
	return lingo.ExpressionFunc(func(d lingo.Dialect) (sql.Data, error) {
		return sql.String(s), nil
	})
}

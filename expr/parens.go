package expr

import (
	"errors"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/sql"
)

// NewParens adds parenthesis around the expression. No spaces are added.
func NewParens(exp lingo.Expression) Parens {
	return Parens{
		exp: exp,
	}
}

type Parens struct {
	exp lingo.Expression
}

func (p Parens) ToSQL(d lingo.Dialect) (sql.Data, error) {
	if check.IsValueNilOrEmpty(p.exp) {
		return nil, errors.New("paren exp cannot be empty")
	}
	s, err := p.exp.ToSQL(d)
	if err != nil {
		return nil, err
	}
	return sql.Surround("(", ")", s), nil
}

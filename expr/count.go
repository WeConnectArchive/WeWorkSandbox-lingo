package expr

import (
	"errors"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/sql"
)

func Count(countOn lingo.Expression) lingo.Expression {
	return &count{
		countOn: countOn,
	}
}

type count struct {
	countOn lingo.Expression
}

func (c count) ToSQL(d lingo.Dialect) (sql.Data, error) {
	if check.IsValueNilOrBlank(c.countOn) {
		return nil, errors.New("countOn value cannot be empty")
	}

	countOn, countOnErr := c.countOn.ToSQL(d)
	if countOnErr != nil {
		return nil, countOnErr
	}

	return sql.String("COUNT").SurroundAppend("(", ")", countOn), nil
}

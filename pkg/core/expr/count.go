package expr

import (
	"errors"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func Count(countOn core.Expression) core.Expression {
	return &count{
		countOn: countOn,
	}
}

type count struct {
	countOn core.Expression
}

func (c count) ToSQL(d core.Dialect) (sql.Data, error) {
	if check.IsValueNilOrBlank(c.countOn) {
		return nil, errors.New("countOn value cannot be empty")
	}

	countOn, countOnErr := c.countOn.ToSQL(d)
	if countOnErr != nil {
		return nil, countOnErr
	}

	return sql.String("COUNT").SurroundAppend("(", ")", countOn), nil
}

package expr

import (
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/sql"
)

// ToList wraps the expressions into a single List. Each expression will be separated by a comma and a space `, `, not
// including the last entry. If there is zero or one expressions, no comma+space is added.
func ToList(exps []lingo.Expression) List {
	return List{
		exps: exps,
	}
}

type List struct {
	exps []lingo.Expression
}

func (l List) ToSQL(d lingo.Dialect) (sql.Data, error) {
	var sqlData = make([]sql.Data, len(l.exps))
	for idx := range l.exps {
		if check.IsValueNilOrEmpty(l.exps) {
			return nil, fmt.Errorf("exps[%d] of list cannot be empty", idx)
		}

		data, err := l.exps[idx].ToSQL(d)
		if err != nil {
			return nil, err
		}
		sqlData[idx] = data
	}
	return sql.Join(", ", sqlData), nil
}

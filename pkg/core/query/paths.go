package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func ExpandTables(paths []core.Expression) []core.Expression {
	var expanded = make([]core.Expression, 0)
	for _, singlePath := range paths {
		if entity, ok := singlePath.(core.Table); ok {
			for _, column := range entity.GetColumns() {
				expanded = append(expanded, column)
			}
		} else {
			expanded = append(expanded, singlePath)
		}
	}
	return expanded
}

// JoinToSQL will call ToSQL on each exp, returning the error if one occurs, and then joins the SQL together with sep
// in between each sql.Data.
func JoinToSQL(d core.Dialect, sep string, exp []core.Expression) (sql.Data, error) {
	// TODO - explore if we could (or really should) pull some of the sqlStr.Data logic into here, or this ToSQL / error
	//  checking logic into the sqlStr.Data logic. Right now, using sqlStr.Join keeps the the copying / appending in one
	//  place, and efficient. However, if broken out, either by using a func closure or something else, then we only
	//  need to loop over each expr once. Currently, it happens twice, once here to generate the SQL, and once
	//  to join the SQL.

	var sqlData = make([]sql.Data, len(exp))
	for idx := range exp {
		data, err := exp[idx].ToSQL(d)
		if err != nil {
			// Join all the data up until this point (slice idx) to produce a somewhat meaningful error message.
			s := sql.Join(sep, sqlData[:idx])
			return nil, NewErrAroundSQL(s, err)
		}
		sqlData[idx] = data
	}
	return sql.Join(sep, sqlData), nil
}

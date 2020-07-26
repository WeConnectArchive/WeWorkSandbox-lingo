package query

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func ExpandTables(paths []lingo.Expression) []lingo.Expression {
	var expanded = make([]lingo.Expression, 0)
	for _, singlePath := range paths {
		if entity, ok := singlePath.(lingo.Table); ok {
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
func JoinToSQL(d lingo.Dialect, sep string, exp []lingo.Expression) (sql.Data, error) {
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

// buildIfNotEmpty will call ToSQL to return the data else the error. If exp is nil or empty, it returns an empty sql
// and no error.
func buildIfNotEmpty(d lingo.Dialect, exp lingo.Expression) (sql.Data, error) {
	if exp == nil {
		return sql.Empty(), nil
	}
	return exp.ToSQL(d)
}

type appendFunc func(prev, new lingo.Expression) expr.Operation

// appendWith takes a previous expression, and will append each newExps using appendFunc to previousExp. If previousExp
// is nil or empty, each newExps will be append to each other using appendFunc - not including previousExp in it. The
// result is a single lingo.Expression from chaining the given expressions.
func appendWith(previousExp lingo.Expression, newExps []lingo.Expression, reduce appendFunc) lingo.Expression {
	for _, exp := range newExps {
		if previousExp == nil {
			previousExp = exp
			continue
		}
		previousExp = reduce(previousExp, exp)
	}
	return previousExp
}

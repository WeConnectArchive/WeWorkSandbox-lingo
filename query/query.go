// Package query is used to build SQL queries with Lingo tables / columns.
package query

import (
	"github.com/weworksandbox/lingo/sql"
)

const (
	// sepPathComma is useful for a separator between SQLs. Ex: column1, column2, column3
	sepPathComma = ", "
	sepSpace     = " "
)

var sqlWhere = sql.String("WHERE")

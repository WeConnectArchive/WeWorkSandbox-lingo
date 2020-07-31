package query

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

// NewNamedOnlyColumn creates a `lingo.Column` of which only the name of the column is filled out.
// Thus, when `ToSQL()` is called, only a single SQL with the value of `name` is returned.
func NewNamedOnlyColumn(name lingo.Expression, parent string) lingo.Column {
	return &stringColumn{name: name, parent: stringParent{name: parent}}
}

type stringColumn struct {
	name   lingo.Expression
	parent stringParent
}

func (s stringColumn) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return s.name.ToSQL(d)
}
func (s stringColumn) GetName() lingo.Expression { return s.name }
func (s stringColumn) GetParent() lingo.Table    { return s.parent }
func (stringColumn) GetAlias() string            { return "" }

type stringParent struct {
	name string
}

func (s stringParent) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return expr.Table(s).ToSQL(d)
}

func (stringParent) GetAlias() lingo.Expression     { return expr.Lit("") }
func (s stringParent) GetTableName() string         { return s.name }
func (s stringParent) GetName() lingo.Expression    { return expr.Lit(s.GetTableName()) }
func (stringParent) GetColumns() []lingo.Expression { return nil }
func (stringParent) GetParent() string              { return "" }

func convertToStringColumns(columns []lingo.Expression) []lingo.Expression {
	if check.IsValueNilOrBlank(columns) {
		return nil
	}

	var expressions = make([]lingo.Expression, 0, len(columns))
	for _, exp := range columns {
		if col, ok := exp.(lingo.Column); ok {
			// TODO we might not even need this entire function or file... maybe remove?
			exp = col.GetName()
		}
		expressions = append(expressions, exp)
	}
	return expressions
}

package query

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

// NewNamedOnlyColumn creates a `core.Column` of which only the name of the column is filled out.
// Thus, when `GetSQL()` is called, only a single SQL with the value of `name` is returned.
func NewNamedOnlyColumn(name, parent string) core.Column {
	return &stringColumn{name: name, parent: stringParent{name: parent}}
}

type stringColumn struct {
	name   string
	parent stringParent
}

func (s stringColumn) GetSQL(_ core.Dialect) (core.SQL, error) {
	return core.NewSQL(s.GetName(), nil), nil
}
func (s stringColumn) GetName() string       { return s.name }
func (s stringColumn) GetParent() core.Table { return s.parent }
func (stringColumn) GetAlias() string        { return "" }

type stringParent struct {
	name string
}

func (s stringParent) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, s)
}

func (stringParent) GetAlias() string          { return "" }
func (s stringParent) GetName() string         { return s.name }
func (stringParent) GetColumns() []core.Column { return []core.Column{} }
func (stringParent) GetParent() string         { return "" }

func convertToStringColumns(columns []core.Column) []core.Expression {
	if helpers.IsValueNilOrBlank(columns) {
		return nil
	}

	var expressions = make([]core.Expression, 0, len(columns))
	for _, column := range columns {
		if helpers.IsValueNilOrBlank(column) {
			return nil
		}
		// TODO we might not even need this entire function or file... maybe remove?
		expressions = append(expressions, NewNamedOnlyColumn(column.GetName(), column.GetParent().GetName()))
	}
	return expressions
}

package dialect

import (
	"fmt"
	"strings"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/join"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sort"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

// NewDefault takes options to configure a Default schema
func NewDefault(opts ...Option) (Default, error) {
	var o options
	for idx := range opts {
		if err := opts[idx](&o); err != nil {
			return Default{}, fmt.Errorf("unable to create default dialect: %w", err)
		}
	}
	return Default{
		includeSchemaName: o.includeSchemaName,
	}, nil
}

// Default schema uses the generic schema methods to work as a basic ANSI schema.
type Default struct {
	includeSchemaName bool
}

func (Default) GetName() string {
	return "Default"
}

func (Default) ValueFormat(count int) sql.Data {
	s := strings.Repeat("?, ", count)
	s = strings.TrimSuffix(s, ", ")
	return sql.String(s)
}

func (Default) SetValueFormat() string {
	return "="
}

func (d Default) ExpandTable(table core.Table) (sql.Data, error) {
	if d.includeSchemaName {
		return ExpandTableWithSchema(table)
	}
	return ExpandTable(table)
}

func (Default) ExpandColumn(column core.Column) (sql.Data, error) {
	return ExpandColumnWithParent(column)
}

func (Default) Operator(left sql.Data, op operator.Operand, values []sql.Data) (sql.Data, error) {
	// No special operations needed beyond ANSI SQL
	return Operator(left, op, values)
}

func (d Default) Value(value []interface{}) (sql.Data, error) {
	return Value(d, value)
}

func (Default) Join(left sql.Data, joinType join.Type, on sql.Data) (sql.Data, error) {
	return Join(left, genericJoinTypeToStr[joinType], on)
}

func (Default) OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error) {
	return OrderBy(left, direction)
}

func (d Default) Set(left sql.Data, value sql.Data) (sql.Data, error) {
	return Set(d, left, value)
}

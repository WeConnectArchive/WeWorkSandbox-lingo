package dialect

import (
	"fmt"
	"strings"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/join"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sort"
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

func (Default) ValueFormat(count int) core.SQL {
	s := strings.Repeat("?, ", count)
	s = strings.TrimSuffix(s, ", ")
	return core.NewSQL(s, nil)
}

func (Default) SetValueFormat() string {
	return "="
}

func (d Default) ExpandTable(table core.Table) (core.SQL, error) {
	if d.includeSchemaName {
		return ExpandTableWithSchema(table)
	}
	return ExpandTable(table)
}

func (Default) ExpandColumn(column core.Column) (core.SQL, error) {
	return ExpandColumnWithParent(column)
}

func (Default) Operator(left core.SQL, op operator.Operand, values []core.SQL) (core.SQL, error) {
	// No special operations needed beyond ANSI SQL
	return Operator(left, op, values)
}

func (d Default) Value(value []interface{}) (core.SQL, error) {
	return Value(d, value)
}

func (Default) Join(left core.SQL, joinType join.Type, on core.SQL) (core.SQL, error) {
	return Join(left, genericJoinTypeToStr[joinType], on)
}

func (Default) OrderBy(left core.SQL, direction sort.Direction) (core.SQL, error) {
	return OrderBy(left, direction)
}

func (d Default) Set(left core.SQL, value core.SQL) (core.SQL, error) {
	return Set(d, left, value)
}

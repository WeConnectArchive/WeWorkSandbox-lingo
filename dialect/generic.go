package dialect

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr/join"
	"github.com/weworksandbox/lingo/expr/sort"
	"github.com/weworksandbox/lingo/sql"
)

// AliasElseName will use the lingo.Alias if non-empty, else the Name is used.
func AliasElseName(n lingo.Name) sql.Data {
	alias, ok := n.(lingo.Alias)
	if aliasStr := alias.GetAlias(); ok && aliasStr != "" {
		return sql.String(aliasStr)
	}
	return sql.String(n.GetName())
}

func ExpandTable(entity lingo.Table) (sql.Data, error) {
	s := sql.String(entity.GetName())
	if alias := entity.GetAlias(); alias != "" {
		s = s.Append(sql.Format(" AS %s", alias))
	}
	return s, nil
}

func ExpandTableWithSchema(entity lingo.Table) (sql.Data, error) {
	s, err := ExpandTable(entity)
	if err != nil {
		return nil, fmt.Errorf("unable to expand table before schema: %w", err)
	}
	return sql.Format("%s.", entity.GetParent()).Append(s), nil
}

func ExpandColumn(column lingo.Column) (sql.Data, error) {
	s := sql.String(column.GetName())
	if a := column.GetAlias(); a != "" {
		s = s.Append(sql.Format(" AS %s", a))
	}
	return s, nil
}

func ExpandColumnWithParent(column lingo.Column) (sql.Data, error) {
	table := AliasElseName(column.GetParent())
	colSQL, err := ExpandColumn(column)
	if err != nil {
		return nil, fmt.Errorf("unable to expand column: %w", err)
	}
	// Append separator prior to column: `table.column`
	return table.Append(sql.String(".").Append(colSQL)), nil
}

type ValueFormatter interface {
	ValueFormat(count int) sql.Data
}

func Value(formatter ValueFormatter, values []interface{}) (sql.Data, error) {
	return formatter.ValueFormat(len(values)).Append(sql.Values(values)), nil
}

var genericJoinTypeToStr = map[join.Type]string{
	join.Inner: "INNER JOIN",
	join.Outer: "OUTER JOIN",
	join.Left:  "LEFT JOIN",
	join.Right: "RIGHT JOIN",
}

func Join(left sql.Data, joinType join.Type, on sql.Data) (sql.Data, error) {
	jTypeStr, ok := genericJoinTypeToStr[joinType]
	if !ok {
		return nil, EnumIsInvalid("join.Type", joinType)
	}

	return sql.String(jTypeStr).
		AppendWithSpace(left).
		AppendWithSpace(sql.String("ON")).
		AppendWithSpace(on), nil
}

type SetFormatter interface {
	SetValueFormat() string
}

func Set(format SetFormatter, left sql.Data, value sql.Data) (sql.Data, error) {
	if check.IsValueNilOrBlank(format) {
		return nil, errors.New("SetFormatter is nil or the interface pointer is nil")
	}
	return left.AppendWithSpace(sql.String(format.SetValueFormat())).AppendWithSpace(value), nil
}

var sortDirectionToStr = map[sort.Direction]string{
	sort.OpAscending:  "ASC",
	sort.OpDescending: "DESC",
}

func OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error) {
	dirStr, ok := sortDirectionToStr[direction]
	if !ok {
		return nil, EnumIsInvalid("sort.Direction", direction)
	}

	return left.AppendWithSpace(sql.String(dirStr)), nil
}

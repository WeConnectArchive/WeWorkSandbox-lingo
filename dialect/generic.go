package dialect

import (
	"errors"

	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr/join"
	"github.com/weworksandbox/lingo/query/sort"
	"github.com/weworksandbox/lingo/sql"
)

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

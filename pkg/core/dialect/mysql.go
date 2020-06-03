package dialect

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/json"
)

// NewMySQL takes options to configure a MySQL schema
func NewMySQL(opts ...Option) (core.Dialect, error) {
	dialect, err := NewDefault(opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to create MySQL dialect: %w", err)
	}
	return MySQL{
		Default: dialect,
	}, nil
}

// MySQL schema has extra MySQL specific features like JSONExtract.
type MySQL struct{ Default }

func (MySQL) GetName() string {
	return "MySQL"
}

func (m MySQL) JSONOperator(left core.SQL, op json.Operand, values []core.SQL) (core.SQL, error) {
	switch op {
	case json.Extract:
		return m.multiPathJSON(left, op, values)
	}

	return nil, expression.ErrorAroundSQL(expression.EnumIsInvalid("json.Operator", op), left.String())
}

func (MySQL) multiPathJSON(left core.SQL, op json.Operand, values []core.SQL) (core.SQL, error) {
	if check.IsValueNilOrBlank(left) {
		return nil, expression.ExpressionIsNil("left")
	}

	opStr, ok := mysqlJSONOperatorToString[op]
	if !ok {
		return nil, expression.EnumIsInvalid("json.Operator", op)
	}

	return core.NewSQLf(opStr).AppendSQL(left.AppendString(", ").CombinePaths(values).SurroundWithParens()), nil
}

var mysqlJSONOperatorToString = map[json.Operand]string{
	json.Extract: "JSON_EXTRACT",
}

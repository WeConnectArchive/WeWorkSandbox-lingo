package dialect

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
	"github.com/weworksandbox/lingo/pkg/core/json"
)

type MySQL struct{ Default }

func (MySQL) GetName() string {
	return "MySQL"
}

func (m MySQL) JSONOperator(left core.SQL, op json.Operand, values []core.SQL) (core.SQL, error) {
	switch op {
	case json.Extract:
		return m.multiPathJson(left, op, values)
	}

	return nil, expression.ErrorAroundSql(expression.EnumIsInvalid("json.Operator", op), left.String())
}

func (MySQL) multiPathJson(left core.SQL, op json.Operand, values []core.SQL) (core.SQL, error) {
	if helpers.IsValueNilOrBlank(left) {
		return nil, expression.ExpressionIsNil("left")
	}

	opStr, ok := mysqlJSONOperatorToString[op]
	if !ok {
		return nil, expression.EnumIsInvalid("json.Operator", op)
	}

	return core.NewSQLf(opStr).AppendSql(left.AppendString(", ").CombinePaths(values).SurroundWithParens()), nil
}

var mysqlJSONOperatorToString = map[json.Operand]string{
	json.Extract: "JSON_EXTRACT",
}

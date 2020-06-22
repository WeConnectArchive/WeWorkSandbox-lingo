package dialect

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/json"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

// NewMySQL takes options to configure a MySQL schema
//
// There are some caveats for MySQL relating to the dialect.Option types.
// If WithSchemaNameIncluded is not true, ensure the database/schema name is in the sql.DB dataSourceName connection
// string. If not, "Error 1046: No database selected" occurs.
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

// Modify will build: LIMIT [offset,]row_count
func (d Default) Modify(m query.Modifier) (sql.Data, error) {
	limit, lWasSet := m.Limit()
	offset, oWasSet := m.Offset()

	if !lWasSet && !oWasSet {
		return sql.Empty(), nil
	}

	s := sql.String("LIMIT ")
	if oWasSet {
		offsetSQL, err := d.Value([]interface{}{offset})
		if err != nil {
			return nil, expression.ErrorAroundSQL(err, s.String())
		}
		s = s.Append(offsetSQL).Append(sql.String(","))
	}
	limitSQL, err := d.Value([]interface{}{limit})
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	return s.Append(limitSQL), nil
}

func (m MySQL) JSONOperator(left sql.Data, op json.Operand, values []sql.Data) (sql.Data, error) {
	switch op {
	case json.Extract:
		return m.multiPathJSON(left, op, values)
	}

	return nil, expression.ErrorAroundSQL(expression.EnumIsInvalid("json.Operator", op), left.String())
}

func (MySQL) multiPathJSON(left sql.Data, op json.Operand, values []sql.Data) (sql.Data, error) {
	if check.IsValueNilOrBlank(left) {
		return nil, expression.ExpressionIsNil("left")
	}

	opStr, ok := mysqlJSONOperatorToString[op]
	if !ok {
		return nil, expression.EnumIsInvalid("json.Operator", op)
	}

	return sql.String(opStr).
		Append(left.Append(sql.String(", ")).
			SurroundAppend("(", ")", sql.Join(", ", values)),
		), nil
}

var mysqlJSONOperatorToString = map[json.Operand]string{
	json.Extract: "JSON_EXTRACT",
}

package path

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/operator"
)

func NewBoolPathWithAlias(e core.Table, name, alias string) BoolPath {
	return BoolPath{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewBoolPath(e core.Table, name string) BoolPath {
	return NewBoolPathWithAlias(e, name, "")
}

type BoolPath struct {
	entity core.Table
	name   string
	alias  string
}

func (b BoolPath) GetParent() core.Table {
	return b.entity
}

func (b BoolPath) GetName() string {
	return b.name
}

func (b BoolPath) GetAlias() string {
	return b.alias
}

func (b BoolPath) As(alias string) BoolPath {
	b.alias = alias
	return b
}

func (b BoolPath) GetSQL(d core.Dialect, sql core.SQL) error {
	return ExpandColumnWithDialect(d, b, sql)
}

func (b BoolPath) To(value bool) core.Set {
	return expression.NewSet(b, expression.NewValue(value))
}

func (b BoolPath) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(b, setExp)
}

func (b BoolPath) Eq(equalTo bool) core.ComboExpression {
	return expression.NewOperator(b, operator.Eq, expression.NewValue(equalTo))
}

func (b BoolPath) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.Eq, equalTo)
}

func (b BoolPath) NotEq(notEqualTo bool) core.ComboExpression {
	return expression.NewOperator(b, operator.NotEq, expression.NewValue(notEqualTo))
}

func (b BoolPath) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.NotEq, notEqualTo)
}

func (b BoolPath) IsNull() core.ComboExpression {
	return expression.NewOperator(b, operator.Null)
}

func (b BoolPath) IsNotNull() core.ComboExpression {
	return expression.NewOperator(b, operator.NotNull)
}

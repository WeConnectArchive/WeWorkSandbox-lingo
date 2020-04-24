package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewBoolPathWithAlias(e core.Table, name, alias string) Bool {
	return Bool{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewBoolPath(e core.Table, name string) Bool {
	return NewBoolPathWithAlias(e, name, "")
}

type Bool struct {
	entity core.Table
	name   string
	alias  string
}

func (b Bool) GetParent() core.Table {
	return b.entity
}

func (b Bool) GetName() string {
	return b.name
}

func (b Bool) GetAlias() string {
	return b.alias
}

func (b Bool) As(alias string) Bool {
	b.alias = alias
	return b
}

func (b Bool) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, b)
}

func (b Bool) To(value bool) core.Set {
	return expression.NewSet(b, expression.NewValue(value))
}

func (b Bool) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(b, setExp)
}

func (b Bool) Eq(equalTo bool) core.ComboExpression {
	return expression.NewOperator(b, operator.Eq, expression.NewValue(equalTo))
}

func (b Bool) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.Eq, equalTo)
}

func (b Bool) NotEq(notEqualTo bool) core.ComboExpression {
	return expression.NewOperator(b, operator.NotEq, expression.NewValue(notEqualTo))
}

func (b Bool) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.NotEq, notEqualTo)
}

func (b Bool) IsNull() core.ComboExpression {
	return expression.NewOperator(b, operator.Null)
}

func (b Bool) IsNotNull() core.ComboExpression {
	return expression.NewOperator(b, operator.NotNull)
}

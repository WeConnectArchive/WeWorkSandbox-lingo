package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewBoolWithAlias(e lingo.Table, name, alias string) Bool {
	return Bool{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewBool(e lingo.Table, name string) Bool {
	return NewBoolWithAlias(e, name, "")
}

type Bool struct {
	entity lingo.Table
	name   string
	alias  string
}

func (b Bool) GetParent() lingo.Table {
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

func (b Bool) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, b)
}

func (b Bool) To(value bool) set.Set {
	return set.NewSet(b, expr.NewValue(value))
}

func (b Bool) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(b, setExp)
}

func (b Bool) Eq(equalTo bool) operator.Binary {
	return operator.NewBinary(b, operator.Eq, expr.NewValue(equalTo))
}

func (b Bool) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.Eq, equalTo)
}

func (b Bool) NotEq(notEqualTo bool) operator.Binary {
	return operator.NewBinary(b, operator.NotEq, expr.NewValue(notEqualTo))
}

func (b Bool) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.NotEq, notEqualTo)
}

func (b Bool) IsNull() operator.Unary {
	return operator.NewUnary(b, operator.Null)
}

func (b Bool) IsNotNull() operator.Unary {
	return operator.NewUnary(b, operator.NotNull)
}

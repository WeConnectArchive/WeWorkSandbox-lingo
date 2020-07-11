package join

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

type Dialect interface {
	Join(left sql.Data, joinType Type, on sql.Data) (sql.Data, error)
}

func NewJoinOn(left lingo.Expression, joinType Type, on lingo.Expression) On {
	return On{
		left:     left,
		on:       on,
		joinType: joinType,
	}
}

type On struct {
	left     lingo.Expression
	on       lingo.Expression
	joinType Type
}

func (j On) And(exp lingo.Expression) lingo.ComboExpression {
	return operator.NewBinary(j, operator.And, exp)
}

func (j On) Or(exp lingo.Expression) lingo.ComboExpression {
	return operator.NewBinary(j, operator.Or, exp)
}

func (j On) ToSQL(d lingo.Dialect) (sql.Data, error) {
	joiner, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'join.Dialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(j.on) {
		return nil, errors.New("join 'on' cannot be empty")
	}
	on, oerr := j.on.ToSQL(d)
	if oerr != nil {
		return nil, oerr
	}

	if check.IsValueNilOrEmpty(j.left) {
		return nil, errors.New("left of join cannot be empty")
	}
	left, lerr := j.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	return joiner.Join(left, j.joinType, on)
}

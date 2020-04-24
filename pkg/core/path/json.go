package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/json"
)

func NewJSONPathWithAlias(e core.Table, name, alias string) JSON {
	return JSON{
		String: NewStringPathWithAlias(e, name, alias),
	}
}

func NewJSONPath(e core.Table, name string) JSON {
	return NewJSONPathWithAlias(e, name, "")
}

type JSON struct {
	String
}

func (j JSON) As(alias string) JSON {
	j.alias = alias
	return j
}

func (j JSON) Extract(paths ...string) core.ComboExpression {
	return expression.NewJSONOperation(j, json.Extract, expression.NewValue(paths))
}

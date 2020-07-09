package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/json"
)

func NewJSONWithAlias(e lingo.Table, name, alias string) JSON {
	return JSON{
		String: NewStringWithAlias(e, name, alias),
	}
}

func NewJSON(e lingo.Table, name string) JSON {
	return NewJSONWithAlias(e, name, "")
}

type JSON struct {
	String
}

func (j JSON) As(alias string) JSON {
	j.alias = alias
	return j
}

func (j JSON) Extract(paths ...string) lingo.ComboExpression {
	return json.NewJSONOperation(j, json.Extract, expr.NewValue(paths))
}

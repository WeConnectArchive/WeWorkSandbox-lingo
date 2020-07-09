package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/json"
)

func NewJSONWithAlias(e core.Table, name, alias string) JSON {
	return JSON{
		String: NewStringWithAlias(e, name, alias),
	}
}

func NewJSON(e core.Table, name string) JSON {
	return NewJSONWithAlias(e, name, "")
}

type JSON struct {
	String
}

func (j JSON) As(alias string) JSON {
	j.alias = alias
	return j
}

func (j JSON) Extract(paths ...string) core.ComboExpression {
	return json.NewJSONOperation(j, json.Extract, expr.NewValue(paths))
}

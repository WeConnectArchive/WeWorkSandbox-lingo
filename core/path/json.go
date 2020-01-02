package path

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/json"
)

func NewJSONPathWithAlias(e core.Table, name, alias string) JSONPath {
	return JSONPath{
		StringPath: NewStringPathWithAlias(e, name, alias),
	}
}

func NewJSONPath(e core.Table, name string) JSONPath {
	return NewJSONPathWithAlias(e, name, "")
}

type JSONPath struct {
	StringPath
}

func (j JSONPath) As(alias string) JSONPath {
	j.alias = alias
	return j
}

func (j JSONPath) Extract(paths ...string) core.ComboExpression {
	return expression.NewJSONOperation(j, json.Extract, expression.NewValues(paths)...)
}

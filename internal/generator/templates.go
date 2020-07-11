package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"
)

var pathTemplate *template.Template
var tableTemplate *template.Template
var exportedTemplate *template.Template
var schemaTemplate *template.Template

func init() {
	var err error
	pathTemplate, err = template.New("path").Parse(pathTemplateString)
	if err != nil {
		panic(fmt.Errorf("unable to create path template: %w", err).Error())
	}
	tableTemplate, err = template.New("table").Parse(tableTemplateString)
	if err != nil {
		panic(fmt.Errorf("unable to create table template: %w", err).Error())
	}
	exportedTemplate, err = template.New("exported").Parse(exportedTemplateString)
	if err != nil {
		panic(fmt.Errorf("unable to create exported template: %w", err).Error())
	}
	schemaTemplate, err = template.New("schema").Parse(schemaTemplateString)
	if err != nil {
		panic(fmt.Errorf("unable to create schema template: %w", err).Error())
	}
}

func generateFromTemplate(t *template.Template, data interface{}) (io.Reader, error) {
	var b strings.Builder
	err := t.Execute(&b, data)
	if err != nil {
		return nil, fmt.Errorf("unable to generate data from template: %w", err)
	}

	formatted, err := format.Source([]byte(b.String()))
	if err != nil {
		return nil, fmt.Errorf("unable to format code after templatizing: %s\n%s", err, b.String())
	}
	return bytes.NewReader(formatted), nil
}

const imports = `
{{ $length := len .Imports }}{{ if gt $length 0 }}
import ({{ range $import := .Imports }}{{ if $importLen := len $import }}
	"{{ $import }}"{{ else }}
{{ end }}{{ end -}}
){{ end }}`

const pathTemplateString = GenPathFileHeader + `

package path
` + imports + `

func New{{ .Name }}WithAlias(e lingo.Table, name, alias string) {{ .Name }} {
	return {{ .Name }}{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func New{{ .Name }}(e lingo.Table, name string) {{ .Name }} {
	return New{{ .Name }}WithAlias(e, name, "")
}

type {{ .Name }} struct {
	entity lingo.Table
	name   string
	alias  string
}

func (p {{ .Name }}) GetParent() lingo.Table {
	return p.entity
}

func (p {{ .Name }}) GetName() string {
	return p.name
}

func (p {{ .Name }}) GetAlias() string {
	return p.alias
}

func (p {{ .Name }}) As(alias string) {{ .Name }} {
	p.alias = alias
	return p
}

func (p {{ .Name }}) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, p)
}

{{- if .Operations.Set }}

func (p {{ .Name }}) To(value {{ .GoType }}) set.Set {
	return set.NewSet(p, expr.NewValue(value))
}

func (p {{ .Name }}) ToExpr(exp lingo.Expression) set.Set {
	return set.NewSet(p, exp)
}
{{- end -}}

{{ if .Operations.Equality }}

func (p {{ .Name }}) Eq(value {{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.Eq, expr.NewValue(value))
}

func (p {{ .Name }}) EqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.Eq, exp)
}

func (p {{ .Name }}) NotEq(value {{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.NotEq, expr.NewValue(value))
}

func (p {{ .Name }}) NotEqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.NotEq, exp)
}
{{- end -}}

{{ if .Operations.Comparison }}

func (p {{ .Name }}) LT(value {{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.LessThan, expr.NewValue(value))
}

func (p {{ .Name }}) LTPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.LessThan, exp)
}

func (p {{ .Name }}) LTOrEq(value {{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.LessThanOrEqual, expr.NewValue(value))
}

func (p {{ .Name }}) LTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.LessThanOrEqual, exp)
}

func (p {{ .Name }}) GT(value {{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThan, expr.NewValue(value))
}

func (p {{ .Name }}) GTPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThan, exp)
}

func (p {{ .Name }}) GTOrEq(value {{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThanOrEqual, expr.NewValue(value))
}

func (p {{ .Name }}) GTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThanOrEqual, exp)
}
{{- end -}}

{{ if .Operations.Nullable }}

func (p {{ .Name }}) IsNull() operator.Unary {
	return operator.NewUnary(p, operator.Null)
}

func (p {{ .Name }}) IsNotNull() operator.Unary {
	return operator.NewUnary(p, operator.NotNull)
}
{{- end -}}

{{ if .Operations.In }}

func (p {{ .Name }}) In(values ...{{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (p {{ .Name }}) InPaths(exps ...lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.In, expr.NewParens(expr.ToList(exps)))
}

func (p {{ .Name }}) NotIn(values ...{{ .GoType }}) operator.Binary {
	return operator.NewBinary(p, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (p {{ .Name }}) NotInPaths(exps ...lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.NotIn, expr.NewParens(expr.ToList(exps)))
}
{{- end -}}

{{ if .Operations.Between }}

func (p {{ .Name }}) Between(first, second {{ .GoType }}) operator.Binary {
	and := expr.NewParens(expr.NewValue(first).And(expr.NewValue(second)))
	return operator.NewBinary(p, operator.Between, and)
}

func (p {{ .Name }}) BetweenPaths(firstExp, secondExp lingo.Expression) operator.Binary {
	and := expr.NewParens(operator.NewBinary(firstExp, operator.And, secondExp))
	return operator.NewBinary(p, operator.Between, and)
}

func (p {{ .Name }}) NotBetween(first, second {{ .GoType }}) operator.Binary {
	and := expr.NewParens(expr.NewValue(first).And(expr.NewValue(second)))
	return operator.NewBinary(p, operator.NotBetween, and)
}

func (p {{ .Name }}) NotBetweenPaths(firstExp, secondExp lingo.Expression) operator.Binary {
	and := expr.NewParens(operator.NewBinary(firstExp, operator.And, secondExp))
	return operator.NewBinary(p, operator.NotBetween, and)
}

{{- end -}}
`

const tableTemplateString = `{{ .GeneratedComment }}

// +build !nolingo

package {{ .PackageName }}

` + imports + `

func As(alias string) {{ .StructName }} {
	return new{{ .StructName }}(alias)
}

func New() {{ .StructName }} {
	return new{{ .StructName }}("")
}

func new{{ .StructName }}(alias string) {{ .StructName }} {
	q := {{ .StructName }}{
		_alias: alias,
	}
	{{- range .Columns }}
	{{ printf "q.%s = %s.New%s(q, \"%s\")" .FieldName .PathType.ShortPkg .PathType.Type .DBName }}
	{{- end }}
	return q
}

type {{ .StructName }} struct {
	_alias     string

	{{ range .Columns -}}
	{{ printf "%s %s.%s" .FieldName .PathType.ShortPkg .PathType.Type }}
	{{ end -}}
}


// lingo.Table Functions

func (q {{ .StructName }}) GetColumns() []lingo.Column {
	return []lingo.Column{
	{{ range .Columns -}}
		{{ printf "q.%s," .FieldName }}
	{{ end -}}
	}
}

func (q {{ .StructName }}) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q {{ .StructName }}) GetAlias() string {
	return q._alias
}

func (q {{ .StructName }}) GetName() string {
	return "{{ .DBName }}"
}

func (q {{ .StructName }}) GetParent() string {
	return "{{ .SchemaName }}"
}

// Column Functions

{{ $receiverName := .StructName }}
{{- range .Columns -}}
func (q {{ $receiverName }}) {{ .MethodName }}() {{ .PathType.ShortPkg }}.{{ .PathType.Type }} {
	return q.{{ .FieldName }}
}

{{ end }}

`

const exportedTemplateString = `{{ .GeneratedComment }}

// +build !nolingo

package {{ .PackageName }}

` + imports + `

var instance = New()

func Q() {{ .StructName }} {
	return instance
}

{{ range .Columns }}
func {{ .MethodName }}() {{ .PathType.ShortPkg }}.{{ .PathType.Type }} {
	return instance.{{ .FieldName }}
}
{{ end }}
`

const schemaTemplateString = `{{ .GeneratedComment }}

// +build !nolingo

package {{ .PackageName }}

` + imports + `

var instance = schema{_name: "{{ .DBName }}"}

func GetSchema() lingo.Name {
	return instance
}

type schema struct {
	_name string
}

func (s schema) GetName() string {
	return s._name
}
`

package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"

	"github.com/weworksandbox/lingo/internal/generate/paths"
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

func FromTemplate(t *template.Template, data interface{}) (io.Reader, error) {
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

const tableTemplateString = `{{ .GeneratedComment }}

// +build !nolingo

package {{ .PackageName }}

{{ template "imports" }}

func As(alias string) {{ .StructName }} {
	return new{{ .StructName }}(alias)
}

func New() {{ .StructName }} {
	return new{{ .StructName }}("")
}

func new{{ .StructName }}(alias string) {{ .StructName }} {
	t := {{ .StructName }}{
		_alias: alias,
	}
	{{- range .Columns }}
	{{ printf "t.%s = %s.New%s(t, \"%s\")" .FieldName .PathType.ShortPkg .PathType.Type .DBName }}
	{{- end }}
	return t
}

type {{ .StructName }} struct {
	_alias     string

	{{ range .Columns -}}
	{{ printf "%s %s.%s" .FieldName .PathType.ShortPkg .PathType.Type }}
	{{ end -}}
}


// lingo.Table Functions

func (t {{ .StructName }}) GetColumns() []lingo.Column {
	return []lingo.Column{
	{{ range .Columns -}}
		{{ printf "t.%s," .FieldName }}
	{{ end -}}
	}
}

func (t {{ .StructName }}) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t {{ .StructName }}) GetAlias() string {
	return t._alias
}

func (t {{ .StructName }}) GetName() string {
	return "{{ .DBName }}"
}

func (t {{ .StructName }}) GetParent() string {
	return "{{ .SchemaName }}"
}

// Column Functions

{{ $receiverName := .StructName }}
{{- range .Columns -}}
func (t {{ $receiverName }}) {{ .MethodName }}() {{ .PathType.ShortPkg }}.{{ .PathType.Type }} {
	return t.{{ .FieldName }}
}

{{ end }}
`

const exportedTemplateString = `{{ .GeneratedComment }}

// +build !nolingo

package {{ .PackageName }}

{{ template "imports" }}

var instance = New()

func {{ .Prefix }}() {{ .StructName }} {
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

{{ template "imports" }}

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

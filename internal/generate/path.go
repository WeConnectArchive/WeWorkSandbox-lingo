package generate

import (
	"io"
	"strings"

	"github.com/weworksandbox/lingo/expr"
)

const (
	GenPathFileHeader = "// Code generated by an internal Lingo tool, genpaths.go - DO NOT EDIT"
)

type Paths struct {
	Package string
	Imports []string
	Paths   []Path
}

func (p Paths) Generate() (io.Reader, error) {
	return FromTemplate(typesTemplate, p)
}

var typesTemplate = createWithImports("types", typesTemplateString)

const typesTemplateString = GenPathFileHeader + `

package {{.Package}}

{{template "imports" .Imports}}

type UnsupportedType struct{}
func (UnsupportedType) ToSQL(_ lingo.Dialect) (sql.Data, error) {
	// TODO - Revisit how we want to deal with unsupported columns. Right now we just ignore them.
	//        Possibly just using the dialect to determine what to do? Dialect options?
	return sql.Empty(), nil
}

{{range $path := .Paths}}

func {{$path.Name}}PtrParam(i *{{$path.GoType}}) {{$path.Name}} {
	return InterfaceParam(i).ToSQL
}

func {{$path.Name}}Param(i {{$path.GoType}}) {{$path.Name}} {
	return InterfaceParam(i).ToSQL
}

type {{$path.Name}} lingo.ExpressionFunc

func (t {{$path.Name}}) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return t(d)
}

{{range $operator, $opInfo := .Operators}}
{{$argCount := len $opInfo.ArgNames}}
{{if gt $argCount 0}}
func (t {{$path.Name}}) {{$operator.String}}({{$opInfo.ToParamName}} {{$path.Name}}) ComboOperation {
	return {{$operator.String}}(t, {{$opInfo.ToParamName}}).ToSQL
}
{{else}}
func (t {{$path.Name}}) {{$operator.String}}() ComboOperation {
	return {{$operator.String}}(t).ToSQL
}
{{- end}}
{{end}}{{end}}
`

type Path struct {
	Name      string
	GoType    string
	Imports   []string
	Operators map[expr.Operator]OperatorInfo
}

type OperatorInfo struct {
	ArgNames []string
}

func (o OperatorInfo) ToParamName() string {
	return strings.Join(o.ArgNames, ", ")
}
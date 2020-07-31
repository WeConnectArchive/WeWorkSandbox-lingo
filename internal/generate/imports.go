package generate

import (
	"text/template"
)

func createWithImports(name string, parse string) *template.Template {
	return createWith(importsTemplate, name, parse)
}

func createWith(prevTemp *template.Template, name string, parse string) *template.Template {
	return template.Must(prevTemp.AddParseTree(name, template.Must(template.New(name).Parse(parse)).Copy()))
}

// Imports creates the import block for Golang and expects . as []strings
var importsTemplate = template.Must(template.New("imports").Parse(importsTemplateString))

const importsTemplateString = `
{{ $length := len . }}{{ if gt $length 0 }}
import ({{ range $import := . }}{{ if $importLen := len $import }}
	"{{ $import }}"{{ else }}
{{ end }}{{ end -}}
){{ end }}
`

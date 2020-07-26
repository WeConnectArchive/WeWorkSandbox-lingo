package generate

import (
	"text/template"
)

// Imports creates the import block for Golang and expects . as []strings
var importsTemplate = template.Must(template.New("imports").Parse(importsTemplateString))
const importsTemplateString = `
{{ $length := len . }}{{ if gt $length 0 }}
import ({{ range $import := . }}{{ if $importLen := len $import }}
	"{{ $import }}"{{ else }}
{{ end }}{{ end -}}
){{ end }}
`

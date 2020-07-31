package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"
)

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

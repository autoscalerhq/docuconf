package docuconf

import (
	"bytes"
	"fmt"
	tmpl "github.com/autoscalerhq/docuconf/internal/template"
	"io"
	"os"
	"text/template"
)

func (c *ConfBuilder) Execute() (string, error) {
	if err := os.MkdirAll(c.outPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("create model pkg path(%s) fail: %s", c.outPath, err)
	}
	var buf bytes.Buffer
	meta := ConfStructMeta{
		Package:    "configuration",
		StructName: "Configuration",
		Fields:     c.options,
	}
	if err := render(tmpl.Struct, &buf, meta); err != nil {
		//goland:noinspection GoPrintFunctions
		return "", fmt.Errorf("failed to render template: %s", err)
	}
	return buf.String(), nil
}

func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

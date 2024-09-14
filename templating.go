package docuconf

import (
	"bytes"
	"fmt"
	tmpl "github.com/autoscalerhq/docuconf/internal/template"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

type OutputFile struct {
	Path    string
	Content string
}

type ExecutionResult struct {
	GoFiles       []OutputFile
	MarkDownFiles []OutputFile
}

func (c *ConfBuilder) Execute() (*ExecutionResult, error) {
	if err := os.MkdirAll(c.outPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("create model pkg path(%s) fail: %s", c.outPath, err)
	}
	var goCodeBuf bytes.Buffer
	meta := ConfStructMeta{
		Package:    "configuration",
		StructName: "Configuration",
		Fields:     c.options,
	}
	if err := render(tmpl.Struct, &goCodeBuf, meta); err != nil {
		//goland:noinspection GoPrintFunctions
		return nil, fmt.Errorf("failed to render template: %s", err)
	}

	var markdownBuf bytes.Buffer

	if err := render(tmpl.MarkDownVariables, &markdownBuf, meta); err != nil {
		return nil, fmt.Errorf("failed to render template: %s", err)
	}
	result := &ExecutionResult{
		GoFiles: []OutputFile{
			{
				Path:    filepath.Join(c.outPath, "configuration.go"),
				Content: goCodeBuf.String(),
			},
		},
		MarkDownFiles: []OutputFile{
			{
				Path:    filepath.Join(c.outPath, "README.md"),
				Content: markdownBuf.String(),
			},
		},
	}
	return result, nil
}

func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

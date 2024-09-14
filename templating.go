package docuconf

import (
	"bytes"
	"fmt"
	tmpl "github.com/autoscalerhq/docuconf/internal/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ConfStructMeta struct {
	Package     string
	StructName  string
	ServiceName string
	Fields      []ConfOptionMeta
}

type ConfOptionMeta struct {
	Name        string
	Type        string
	Description string
	Required    bool
	Default     string
	SharedWith  []string
}

type OutputFile struct {
	Path    string
	Content string
}

type ExecutionResult struct {
	GoFiles       []OutputFile
	MarkDownFiles []OutputFile
}

func convertConfOptionsToMeta(s *Service) []ConfOptionMeta {
	var result []ConfOptionMeta
	for _, o := range s.builder.options {
		result = append(result, ConfOptionMeta{
			Name:        o.Name,
			Type:        o.Type,
			Description: o.Description,
			Required:    o.Required,
			Default:     o.Default,
			SharedWith:  convertServicesToStrings(o.Services, s),
		})
	}
	return result
}

func convertServicesToStrings(services []*Service, currentService *Service) []string {
	var result []string
	for _, service := range services {
		if service != currentService {
			result = append(result, service.Name+", ")
		}
	}
	if len(result) > 0 {
		result[len(result)-1] = strings.TrimRight(result[len(result)-1], ", ")
	}
	return result
}

func WriteAll(services []*Service) error {
	for _, service := range services {
		err := service.Write()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Write() error {
	result, err := s.Execute()
	if err != nil {
		return err
	}
	for _, file := range result.GoFiles {
		err := os.WriteFile(file.Path, []byte(file.Content), 0644)
		if err != nil {
			return err
		}
	}
	for _, file := range result.MarkDownFiles {
		err := os.WriteFile(file.Path, []byte(file.Content), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Execute() (*ExecutionResult, error) {
	if err := os.MkdirAll(s.Path, os.ModePerm); err != nil {
		return nil, fmt.Errorf("create model pkg path(%s) fail: %s", s.Path, err)
	}
	var goCodeBuf bytes.Buffer
	meta := ConfStructMeta{
		Package:     s.Package,
		StructName:  "Configuration",
		ServiceName: s.Name,
		Fields:      convertConfOptionsToMeta(s),
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
				Path:    filepath.Join(s.Path, "configuration.go"),
				Content: goCodeBuf.String(),
			},
		},
		MarkDownFiles: []OutputFile{
			{
				Path:    filepath.Join(s.Path, "README.md"),
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

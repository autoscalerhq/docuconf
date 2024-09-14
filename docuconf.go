package docuconf

import (
	"fmt"
)

type ConfOption struct {
	Name        string
	Type        string
	Description string
	Required    bool
	Default     string
}
type ConfBuilder struct {
	outPath string
	// a map of all the configuration options
	options []ConfOption
}

type ConfStructMeta struct {
	Package    string
	StructName string
	Fields     []ConfOption
}

func NewConfBuilder(outPath string) *ConfBuilder {
	return &ConfBuilder{options: []ConfOption{}, outPath: outPath}
}

func (c *ConfBuilder) AddString(name string, description string, required bool, defaultValue string) *ConfBuilder {
	if len(description) < 5 {
		panic(fmt.Errorf("AddString(%s, %s) Failed: description must be at least 5 characters long. This is to ensure your configuration is well documented", name, description))
	}
	if len(name) == 0 {
		panic(fmt.Errorf("AddString(%s, %s) Failed: name must be a non empty string", name, description))
	}
	c.options = append(c.options, ConfOption{Name: name, Description: description, Type: "string", Required: required, Default: defaultValue})
	return c
}

package template

const Struct = Header + `
import "github.com/autoscalerhq/docuconf"

type {{.StructName}} struct {
  {{- range .Fields}} 

  // {{.Description}}
  // {{if .SharedWith}}also used in: {{ range .SharedWith}}{{.}}{{end}}{{end}}
  {{.Name}} {{.Type}} ` + "`env:\"" + `{{.EnvName}}` + "\"`" + `
  {{- end}}
}

func Load{{.StructName}}(path string) ({{.StructName}}, error) {
	env, err := docuconf.LoadDotEnv(path, New{{.StructName}}())
	return env, err
}

func New{{.StructName}}() {{.StructName}} {
	return {{.StructName}}{
	{{- range .Fields}} 
		{{if .Default}}{{.Name}}: {{.Default}},{{end}}
  	{{- end}}
	}
}

`

const MarkDownVariables = NoEditMark + `

# {{.ServiceName}} Configuration
{{range .Fields}}
---
### {{.Name}} - ` + "`{{.Type}}`" + ` - ({{if .Required}}Required{{else}}Optional{{end}})
{{if .SharedWith}}Also used in: {{ range .SharedWith}}{{.}}{{end}}
{{end}}
**Description**

{{.Description}}
{{end}}
---
`

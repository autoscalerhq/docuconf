package template

const Struct = Header + `
type {{.StructName}} struct {
  {{- range .Fields}} 

  // {{.Description}}
  // {{if .SharedWith}}also used in: {{ range .SharedWith}}{{.}}{{end}}{{end}}
  {{.Name}} {{.Type}} 
  {{- end}}
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

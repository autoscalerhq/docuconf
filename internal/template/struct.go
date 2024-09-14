package template

const Struct = Header + `
type {{.StructName}} struct {
  {{- range .Fields}}
  {{.Name}} {{.Type}} // {{.Description}}
  {{- end}}
}
`

const MarkDownVariables = NoEditMark + `
# Configuration

{{range .Fields}}
## {{.Name}} - {{.Type}} - 
Description: {{.Description}}
{{end}}
`

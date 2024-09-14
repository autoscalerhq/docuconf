package template

const Struct = Header + `
type {{.StructName}} struct {
  {{- range .Fields}}
  {{.Name}} {{.Type}} // {{.Description}}
  {{- end}}
}
`

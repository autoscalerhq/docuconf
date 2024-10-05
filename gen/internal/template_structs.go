package internal

type ConfStructMeta struct {
	Package     string
	StructName  string
	ServiceName string
	Fields      []ConfOptionMeta
}

type ConfOptionMeta struct {
	Name        string
	EnvName     string
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

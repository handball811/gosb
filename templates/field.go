package templates

type Field struct {
	Name           string
	VarName        string
	Type           string
	Pointer        bool
	Optional       bool
	DefaultFunc    string
	ValidationFunc string
}

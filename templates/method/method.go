package method

type Config[T any] struct {
	Method
	Body T
}

type Method struct {
	NameVar  string
	Name     string
	FuncName string
	Args     map[string]string
	Returns  []string
	Body     Body
}

func (m *Method) InlineTemplate() func(string, any) string {
	return InlineTemplate
}

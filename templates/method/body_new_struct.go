package method

var _ Body = new(BodyNewStruct)

type BodyNewStruct struct {
	VarPrefix string
	Fields    []Field
}

func (b *BodyNewStruct) Name() string {
	return "body_new_struct"
}

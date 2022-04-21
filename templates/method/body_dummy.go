package method

var _ Body = new(BodyDummy)

type BodyDummy struct{}

func (b *BodyDummy) Name() string {
	return "body_dummy"
}

package templates

import "github.com/handball811/gosb/templates/parse"

const BodyDummyName = "body_dummy"

var BodyDummyTmpl = parse.GenerateTemplate(
	BodyDummyName,
	`return nil`,
)

var _ Body = new(BodyDummy)

type BodyDummy struct{}

func (b *BodyDummy) Name() string {
	return BodyDummyName
}

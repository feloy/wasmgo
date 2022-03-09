package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewHeader1(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h1",
	}
}
func NewHeader2(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h2",
	}
}
func NewHeader3(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h3",
	}
}
func NewHeader4(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h4",
	}
}
func NewHeader5(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h5",
	}
}
func NewHeader6(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h6",
	}
}

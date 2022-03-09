package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewDiv() *dom.Tag {
	return &dom.Tag{Name: "div"}
}
func NewParagraph(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "p",
	}
}

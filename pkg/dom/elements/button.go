package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewButton(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "button",
	}
}

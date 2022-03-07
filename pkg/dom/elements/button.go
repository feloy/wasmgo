package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewButton(title string) *dom.Tag {
	return &dom.Tag{
		Name:      "button",
		InnerHTML: title,
	}
}

package elements

import "github.com/feloy/wasmgo/pkg/dom"

// NewDiv returns a new empty "<div>" tag
func NewDiv() *dom.Tag {
	return &dom.Tag{
		Name: "div",
	}
}

// NewParagraph returns a new "<p>" tag with text content
func NewParagraph(text string) *dom.Tag {
	return &dom.Tag{
		Name:      "p",
		InnerHTML: text,
	}
}

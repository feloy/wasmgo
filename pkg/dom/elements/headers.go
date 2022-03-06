package elements

import "github.com/feloy/wasmgo/pkg/dom"

// NewHeader1 returns a new "<h1>" tag with title content
func NewHeader1(title string) *dom.Tag {
	return &dom.Tag{
		Name:      "h1",
		InnerHTML: title,
	}
}

// NewHeader2 returns a new "<h2>" tag with title content
func NewHeader2(title string) *dom.Tag {
	return &dom.Tag{
		Name:      "h2",
		InnerHTML: title,
	}
}

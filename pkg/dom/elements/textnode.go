package elements

import "github.com/feloy/wasmgo/pkg/dom"

// NewText returns a new Text node with text content
func NewText(text string) *dom.TextNode {
	return &dom.TextNode{
		Text: text,
	}
}

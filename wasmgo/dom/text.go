package dom

import "syscall/js"

type TextNode struct {
	Text string
}

// NewText returns a new Text node with text content
func NewText(text string) *TextNode {
	return &TextNode{
		Text: text,
	}
}

func (o *TextNode) Render(document js.Value) js.Value {
	return document.Call("createTextNode", o.Text)
}

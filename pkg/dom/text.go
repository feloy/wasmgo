package dom

type TextNode struct {
	Text string
}

// NewText returns a new Text node with text content
func NewText(text string) *TextNode {
	return &TextNode{
		Text: text,
	}
}

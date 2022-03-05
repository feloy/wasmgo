package dom

type TextNode struct {
	Text string
}

func NewText(text string) *TextNode {
	return &TextNode{
		Text: text,
	}
}

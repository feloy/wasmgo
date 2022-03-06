package elements

import "github.com/feloy/wasmgo/pkg/dom"

type HyperlinkOptions struct {
	Destination string
	Relation    string
}

// NewHyperlink returns a new "<a>" tag with text content
func NewHyperlink(text string, options HyperlinkOptions) *dom.Tag {
	return &dom.Tag{
		Name:      "a",
		InnerHTML: text,
		Attributes: map[string]string{
			"href": options.Destination,
			"rel":  options.Relation,
		},
	}
}

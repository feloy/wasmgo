package elements

import "github.com/feloy/wasmgo/pkg/dom"

type HyperlinkOptions struct {
	Destination string
	Relation    string
}

func NewHyperlink(inner string, options HyperlinkOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"href": options.Destination,
			"rel":  options.Relation,
		},
		InnerHTML: inner,
		Name:      "a",
	}
}

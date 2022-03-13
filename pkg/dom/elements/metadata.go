// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewHead() *dom.Tag {
	return &dom.Tag{Name: "head"}
}

func NewTitle(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "title",
	}
}

func NewBase() *dom.Tag {
	return &dom.Tag{
		Name:       "base",
		OmitEndTag: true,
	}
}

func NewLink() *dom.Tag {
	return &dom.Tag{
		Name:       "link",
		OmitEndTag: true,
	}
}

func NewMeta() *dom.Tag {
	return &dom.Tag{
		Name:       "meta",
		OmitEndTag: true,
	}
}

func NewStyle(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "style",
	}
}

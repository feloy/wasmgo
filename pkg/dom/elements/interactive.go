// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

type DetailsOptions struct {
	Open bool
}

func NewDetails(options DetailsOptions) *dom.Tag {
	return &dom.Tag{
		BoolAttributes: map[string]bool{"open": options.Open},
		Name:           "details",
	}
}

func NewSummary(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "summary",
	}
}

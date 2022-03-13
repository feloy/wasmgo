// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

type DetailsOptions struct {
	Open
}

func NewDetails(options DetailsOptions) *dom.Tag {
	return &dom.Tag{Name: "details"}
}

func NewSummary(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "summary",
	}
}

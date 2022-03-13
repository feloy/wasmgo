// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

type InsOptions struct {
	Cite
	Datetime
}

func NewIns(inner string, options InsOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "ins",
	}
}

type DelOptions struct {
	Cite
	Datetime
}

func NewDel(inner string, options DelOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "del",
	}
}

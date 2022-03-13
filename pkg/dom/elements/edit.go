// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

type InsOptions struct {
	Cite     string
	Datetime string
}

func NewIns(inner string, options InsOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"cite":     options.Cite,
			"datetime": options.Datetime,
		},
		InnerHTML: inner,
		Name:      "ins",
	}
}

type DelOptions struct {
	Cite     string
	Datetime string
}

func NewDel(inner string, options DelOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"cite":     options.Cite,
			"datetime": options.Datetime,
		},
		InnerHTML: inner,
		Name:      "del",
	}
}

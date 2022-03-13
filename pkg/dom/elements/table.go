// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewTable() *dom.Tag {
	return &dom.Tag{Name: "table"}
}

func NewCaption(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "caption",
	}
}

type ColgroupOptions struct {
	Span
}

func NewColgroup(options ColgroupOptions) *dom.Tag {
	return &dom.Tag{Name: "colgroup"}
}

type ColOptions struct {
	Span
}

func NewCol(options ColOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "col",
		OmitEndTag: true,
	}
}

func NewTbody() *dom.Tag {
	return &dom.Tag{Name: "tbody"}
}

func NewThead() *dom.Tag {
	return &dom.Tag{Name: "thead"}
}

func NewTfoot() *dom.Tag {
	return &dom.Tag{Name: "tfoot"}
}

func NewTr() *dom.Tag {
	return &dom.Tag{Name: "tr"}
}

type TdOptions struct {
	Colspan
	Rowspan
	Headers
}

func NewTd(inner string, options TdOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "td",
	}
}

type ThOptions struct {
	Colspan
	Rowspan
	Headers
	Scope
	Abbr
}

func NewTh(inner string, options ThOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "th",
	}
}

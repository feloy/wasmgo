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
func NewColgroup() *dom.Tag {
	return &dom.Tag{Name: "colgroup"}
}
func NewCol() *dom.Tag {
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
func NewTd(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "td",
	}
}
func NewTh(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "th",
	}
}

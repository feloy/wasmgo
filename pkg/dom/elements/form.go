// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewForm(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "form",
	}
}
func NewLabel(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "label",
	}
}
func NewInput() *dom.Tag {
	return &dom.Tag{
		Name:       "input",
		OmitEndTag: true,
	}
}
func NewButton(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "button",
	}
}
func NewSelect() *dom.Tag {
	return &dom.Tag{Name: "select"}
}
func NewDatalist(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "datalist",
	}
}
func NewOptgroup() *dom.Tag {
	return &dom.Tag{Name: "optgroup"}
}
func NewOption(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "option",
	}
}
func NewTextarea(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "textarea",
	}
}
func NewOutput(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "output",
	}
}
func NewProgress(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "progress",
	}
}
func NewMeter(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "meter",
	}
}
func NewFieldset(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "fieldset",
	}
}
func NewLegend(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "legend",
	}
}

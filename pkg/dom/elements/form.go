// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewForm() *dom.Tag {
	return &dom.Tag{Name: "form"}
}
func NewLabel() *dom.Tag {
	return &dom.Tag{Name: "label"}
}
func NewInput() *dom.Tag {
	return &dom.Tag{Name: "input"}
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
func NewDatalist() *dom.Tag {
	return &dom.Tag{Name: "datalist"}
}
func NewOptgroup() *dom.Tag {
	return &dom.Tag{Name: "optgroup"}
}
func NewOption() *dom.Tag {
	return &dom.Tag{Name: "option"}
}
func NewTextarea() *dom.Tag {
	return &dom.Tag{Name: "textarea"}
}
func NewOutput() *dom.Tag {
	return &dom.Tag{Name: "output"}
}
func NewProgress() *dom.Tag {
	return &dom.Tag{Name: "progress"}
}
func NewMeter() *dom.Tag {
	return &dom.Tag{Name: "meter"}
}
func NewFieldset() *dom.Tag {
	return &dom.Tag{Name: "fieldset"}
}
func NewLegend() *dom.Tag {
	return &dom.Tag{Name: "legend"}
}

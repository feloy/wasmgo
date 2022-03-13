// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

type FormOptions struct {
	AcceptCharset
	Action
	Autocomplete
	Enctype
	Method
	Name
	Novalidate
	Target
	Rel
}

func NewForm(inner string, options FormOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "form",
	}
}

type LabelOptions struct {
	For
}

func NewLabel(inner string, options LabelOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "label",
	}
}

type InputOptions struct {
	Accept
	Alt
	Autocomplete
	Checked
	Dirname
	Disabled
	Form
	Formaction
	Formenctype
	Formmethod
	Formnovalidate
	Formtarget
	Height
	List
	Max
	Maxlength
	Min
	Minlength
	Multiple
	Name
	Pattern
	Placeholder
	Readonly
	Required
	Size
	Src
	Step
	Type
	Value
	Width
}

func NewInput(options InputOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "input",
		OmitEndTag: true,
	}
}

type ButtonOptions struct {
	Disabled
	Form
	Formaction
	Formenctype
	Formmethod
	Formnovalidate
	Formtarget
	Name
	Type
	Value
}

func NewButton(inner string, options ButtonOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "button",
	}
}

type SelectOptions struct {
	Autocomplete
	Disabled
	Form
	Multiple
	Name
	Required
	Size
}

func NewSelect(options SelectOptions) *dom.Tag {
	return &dom.Tag{Name: "select"}
}

func NewDatalist(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "datalist",
	}
}

type OptgroupOptions struct {
	Disabled
	Label
}

func NewOptgroup(options OptgroupOptions) *dom.Tag {
	return &dom.Tag{Name: "optgroup"}
}

type OptionOptions struct {
	Disabled
	Label
	Selected
	Value
}

func NewOption(inner string, options OptionOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "option",
	}
}

type TextareaOptions struct {
	Autocomplete
	Cols
	Dirname
	Disabled
	Form
	Maxlength
	Minlength
	Name
	Placeholder
	Readonly
	Required
	Rows
	Wrap
}

func NewTextarea(inner string, options TextareaOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "textarea",
	}
}

type OutputOptions struct {
	For
	Form
	Name
}

func NewOutput(inner string, options OutputOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "output",
	}
}

type ProgressOptions struct {
	Value
	Max
}

func NewProgress(inner string, options ProgressOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "progress",
	}
}

type MeterOptions struct {
	Value
	Min
	Max
	Low
	High
	Optimum
}

func NewMeter(inner string, options MeterOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "meter",
	}
}

type FieldsetOptions struct {
	Disabled
	Form
	Name
}

func NewFieldset(inner string, options FieldsetOptions) *dom.Tag {
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

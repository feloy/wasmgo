// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

const (
	Form_Autocomplete_On  = "on"
	Form_Autocomplete_Off = "off"
)

const (
	Form_Enctype_ApplicationXWwwFormUrlencoded = "application/x-www-form-urlencoded"
	Form_Enctype_MultipartFormData             = "multipart/form-data"
	Form_Enctype_TextPlain                     = "text/plain"
)

const (
	Form_Method_Get    = "get"
	Form_Method_Post   = "post"
	Form_Method_Dialog = "dialog"
)

const (
	Form_Rel_Opener     = "opener"
	Form_Rel_Next       = "next"
	Form_Rel_Noreferrer = "noreferrer"
	Form_Rel_Prev       = "prev"
	Form_Rel_External   = "external"
	Form_Rel_Help       = "help"
	Form_Rel_License    = "license"
	Form_Rel_Search     = "search"
	Form_Rel_Noopener   = "noopener"
	Form_Rel_Nofollow   = "nofollow"
)

type FormOptions struct {
	AcceptCharset string
	Action        string
	Autocomplete  string
	Enctype       string
	Method        string
	Name          string
	Novalidate    bool
	Target        string
	Rel           []string
}

func NewForm(inner string, options FormOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"accept-charset": options.Accept - Charset,
			"action":         options.Action,
			"autocomplete":   options.Autocomplete,
			"enctype":        options.Enctype,
			"method":         options.Method,
			"name":           options.Name,
			"target":         options.Target,
		},
		BoolAttributes: map[string]bool{"novalidate": options.Novalidate},
		InnerHTML:      inner,
		Name:           "form",
		SSLAttributes:  map[string][]string{"rel": options.Rel},
	}
}

type LabelOptions struct {
	For string
}

func NewLabel(inner string, options LabelOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{"for": options.For},
		InnerHTML:  inner,
		Name:       "label",
	}
}

const (
	Input_Formenctype_ApplicationXWwwFormUrlencoded = "application/x-www-form-urlencoded"
	Input_Formenctype_MultipartFormData             = "multipart/form-data"
	Input_Formenctype_TextPlain                     = "text/plain"
)

const (
	Input_Formmethod_Get    = "get"
	Input_Formmethod_Post   = "post"
	Input_Formmethod_Dialog = "dialog"
)

const (
	Input_Type_Hidden        = "hidden"
	Input_Type_Text          = "text"
	Input_Type_Search        = "search"
	Input_Type_Tel           = "tel"
	Input_Type_Url           = "url"
	Input_Type_Email         = "email"
	Input_Type_Password      = "password"
	Input_Type_Date          = "date"
	Input_Type_Month         = "month"
	Input_Type_Week          = "week"
	Input_Type_Time          = "time"
	Input_Type_DatetimeLocal = "datetime-local"
	Input_Type_Number        = "number"
	Input_Type_Range         = "range"
	Input_Type_Color         = "color"
	Input_Type_Checkbox      = "checkbox"
	Input_Type_Radio         = "radio"
	Input_Type_File          = "file"
	Input_Type_Submit        = "submit"
	Input_Type_Image         = "image"
	Input_Type_Reset         = "reset"
	Input_Type_Button        = "button"
)

type InputOptions struct {
	Accept         string
	Alt            string
	Autocomplete   string
	Checked        bool
	Dirname        string
	Disabled       bool
	Form           string
	Formaction     string
	Formenctype    string
	Formmethod     string
	Formnovalidate bool
	Formtarget     string
	Height         int
	List           string
	Max            int
	Maxlength      int
	Min            int
	Minlength      int
	Multiple       bool
	Name           string
	Pattern        string
	Placeholder    string
	Readonly       bool
	Required       bool
	Size           int
	Src            string
	Step           string
	Type           string
	Value          string
	Width          int
}

func NewInput(options InputOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"accept":       options.Accept,
			"alt":          options.Alt,
			"autocomplete": options.Autocomplete,
			"dirname":      options.Dirname,
			"form":         options.Form,
			"formaction":   options.Formaction,
			"formenctype":  options.Formenctype,
			"formmethod":   options.Formmethod,
			"formtarget":   options.Formtarget,
			"list":         options.List,
			"name":         options.Name,
			"pattern":      options.Pattern,
			"placeholder":  options.Placeholder,
			"src":          options.Src,
			"step":         options.Step,
			"type":         options.Type,
			"value":        options.Value,
		},
		BoolAttributes: map[string]bool{
			"checked":        options.Checked,
			"disabled":       options.Disabled,
			"formnovalidate": options.Formnovalidate,
			"multiple":       options.Multiple,
			"readonly":       options.Readonly,
			"required":       options.Required,
		},
		IntAttributes: map[string]int{
			"height":    options.Height,
			"max":       options.Max,
			"maxlength": options.Maxlength,
			"min":       options.Min,
			"minlength": options.Minlength,
			"size":      options.Size,
			"width":     options.Width,
		},
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

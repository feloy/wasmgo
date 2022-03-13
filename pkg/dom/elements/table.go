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
	Span int
}

func NewColgroup(options ColgroupOptions) *dom.Tag {
	return &dom.Tag{
		IntAttributes: map[string]int{"span": options.Span},
		Name:          "colgroup",
	}
}

type ColOptions struct {
	Span int
}

func NewCol(options ColOptions) *dom.Tag {
	return &dom.Tag{
		IntAttributes: map[string]int{"span": options.Span},
		Name:          "col",
		OmitEndTag:    true,
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
	Colspan int
	Rowspan int
	Headers []string
}

func NewTd(inner string, options TdOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		IntAttributes: map[string]int{
			"colspan": options.Colspan,
			"rowspan": options.Rowspan,
		},
		Name:          "td",
		SSLAttributes: map[string][]string{"headers": options.Headers},
	}
}

const (
	Th_Scope_Row      = "row"
	Th_Scope_Col      = "col"
	Th_Scope_Rowgroup = "rowgroup"
	Th_Scope_Colgroup = "colgroup"
	Th_Scope_Auto     = "auto"
)

type ThOptions struct {
	Colspan int
	Rowspan int
	Headers []string
	Scope   string
	Abbr    string
}

func NewTh(inner string, options ThOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"abbr":  options.Abbr,
			"scope": options.Scope,
		},
		InnerHTML: inner,
		IntAttributes: map[string]int{
			"colspan": options.Colspan,
			"rowspan": options.Rowspan,
		},
		Name:          "th",
		SSLAttributes: map[string][]string{"headers": options.Headers},
	}
}

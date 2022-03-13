// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewP(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "p",
	}
}

func NewHr() *dom.Tag {
	return &dom.Tag{
		Name:       "hr",
		OmitEndTag: true,
	}
}

func NewPre(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "pre",
	}
}

type BlockquoteOptions struct {
	Cite string
}

func NewBlockquote(inner string, options BlockquoteOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{"cite": options.Cite},
		InnerHTML:  inner,
		Name:       "blockquote",
	}
}

const (
	Ol_Type_1 = "1"
	Ol_Type_a = "a"
	Ol_Type_A = "A"
	Ol_Type_i = "i"
	Ol_Type_I = "I"
)

type OlOptions struct {
	Reversed bool
	Start    int
	Type     string
}

func NewOl(options OlOptions) *dom.Tag {
	return &dom.Tag{
		Attributes:     map[string]string{"type": options.Type},
		BoolAttributes: map[string]bool{"reversed": options.Reversed},
		IntAttributes:  map[string]int{"start": options.Start},
		Name:           "ol",
	}
}

func NewUl() *dom.Tag {
	return &dom.Tag{Name: "ul"}
}

func NewMenu() *dom.Tag {
	return &dom.Tag{Name: "menu"}
}

type LiOptions struct {
	Value int
}

func NewLi(inner string, options LiOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML:     inner,
		IntAttributes: map[string]int{"value": options.Value},
		Name:          "li",
	}
}

func NewDl() *dom.Tag {
	return &dom.Tag{Name: "dl"}
}

func NewDt(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "dt",
	}
}

func NewDd(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "dd",
	}
}

func NewFigure(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "figure",
	}
}

func NewFigcaption(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "figcaption",
	}
}

func NewMain(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "main",
	}
}

func NewDiv(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "div",
	}
}

// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

type AOptions struct {
	Destination string
	Relation    string
}

func NewA(inner string, options AOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"href": options.Destination,
			"rel":  options.Relation,
		},
		InnerHTML: inner,
		Name:      "a",
	}
}
func NewEm(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "em",
	}
}
func NewStrong(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "strong",
	}
}
func NewSmall(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "small",
	}
}
func NewS(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "s",
	}
}
func NewCite(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "cite",
	}
}
func NewQ(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "q",
	}
}
func NewDfn(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "dfn",
	}
}
func NewAbbr(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "abbr",
	}
}
func NewRuby(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "ruby",
	}
}
func NewRt(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "rt",
	}
}
func NewRp(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "rp",
	}
}
func NewData(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "data",
	}
}
func NewTime(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "time",
	}
}
func NewCode(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "code",
	}
}
func NewVar(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "var",
	}
}
func NewSamp(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "samp",
	}
}
func NewKbd(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "kbd",
	}
}
func NewSub(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "sub",
	}
}
func NewSup(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "sup",
	}
}
func NewI(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "i",
	}
}
func NewB(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "b",
	}
}
func NewU(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "u",
	}
}
func NewMark(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "mark",
	}
}
func NewBdi(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "bdi",
	}
}
func NewBdo(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "bdo",
	}
}
func NewSpan(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "span",
	}
}
func NewBr() *dom.Tag {
	return &dom.Tag{
		Name:       "br",
		OmitEndTag: true,
	}
}
func NewWbr() *dom.Tag {
	return &dom.Tag{
		Name:       "wbr",
		OmitEndTag: true,
	}
}

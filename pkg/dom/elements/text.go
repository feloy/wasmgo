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
func NewEm() *dom.Tag {
	return &dom.Tag{Name: "em"}
}
func NewStrong() *dom.Tag {
	return &dom.Tag{Name: "strong"}
}
func NewSmall() *dom.Tag {
	return &dom.Tag{Name: "small"}
}
func NewS() *dom.Tag {
	return &dom.Tag{Name: "s"}
}
func NewCite() *dom.Tag {
	return &dom.Tag{Name: "cite"}
}
func NewQ() *dom.Tag {
	return &dom.Tag{Name: "q"}
}
func NewDfn() *dom.Tag {
	return &dom.Tag{Name: "dfn"}
}
func NewAbbr() *dom.Tag {
	return &dom.Tag{Name: "abbr"}
}
func NewRuby() *dom.Tag {
	return &dom.Tag{Name: "ruby"}
}
func NewRt() *dom.Tag {
	return &dom.Tag{Name: "rt"}
}
func NewRp() *dom.Tag {
	return &dom.Tag{Name: "rp"}
}
func NewData() *dom.Tag {
	return &dom.Tag{Name: "data"}
}
func NewTime() *dom.Tag {
	return &dom.Tag{Name: "time"}
}
func NewCode() *dom.Tag {
	return &dom.Tag{Name: "code"}
}
func NewVar() *dom.Tag {
	return &dom.Tag{Name: "var"}
}
func NewSamp() *dom.Tag {
	return &dom.Tag{Name: "samp"}
}
func NewKbd() *dom.Tag {
	return &dom.Tag{Name: "kbd"}
}
func NewSub() *dom.Tag {
	return &dom.Tag{Name: "sub"}
}
func NewSup() *dom.Tag {
	return &dom.Tag{Name: "sup"}
}
func NewI() *dom.Tag {
	return &dom.Tag{Name: "i"}
}
func NewB() *dom.Tag {
	return &dom.Tag{Name: "b"}
}
func NewU() *dom.Tag {
	return &dom.Tag{Name: "u"}
}
func NewMark() *dom.Tag {
	return &dom.Tag{Name: "mark"}
}
func NewBdi() *dom.Tag {
	return &dom.Tag{Name: "bdi"}
}
func NewBdo() *dom.Tag {
	return &dom.Tag{Name: "bdo"}
}
func NewSpan() *dom.Tag {
	return &dom.Tag{Name: "span"}
}
func NewBr() *dom.Tag {
	return &dom.Tag{Name: "br"}
}
func NewWbr() *dom.Tag {
	return &dom.Tag{Name: "wbr"}
}

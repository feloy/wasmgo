// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

const (
	A_Rel_Alternate  = "alternate"
	A_Rel_Author     = "author"
	A_Rel_Bookmark   = "bookmark"
	A_Rel_External   = "external"
	A_Rel_Help       = "help"
	A_Rel_License    = "license"
	A_Rel_Next       = "next"
	A_Rel_Nofollow   = "nofollow"
	A_Rel_Noopener   = "noopener"
	A_Rel_Noreferrer = "noreferrer"
	A_Rel_Opener     = "opener"
	A_Rel_Prev       = "prev"
	A_Rel_Search     = "search"
	A_Rel_Tag        = "tag"
)

const (
	A_Referrerpolicy_NoReferrer                  = "no-referrer"
	A_Referrerpolicy_NoReferrerWhenDowngrade     = "no-referrer-when-downgrade"
	A_Referrerpolicy_SameOrigin                  = "same-origin"
	A_Referrerpolicy_Origin                      = "origin"
	A_Referrerpolicy_StrictOrigin                = "strict-origin"
	A_Referrerpolicy_OriginWhenCrossOrigin       = "origin-when-cross-origin"
	A_Referrerpolicy_StrictOriginWhenCrossOrigin = "strict-origin-when-cross-origin"
	A_Referrerpolicy_UnsafeUrl                   = "unsafe-url"
)

type AOptions struct {
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            []string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

func NewA(inner string, options AOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"download":       options.Download,
			"href":           options.Href,
			"hreflang":       options.Hreflang,
			"ping":           options.Ping,
			"referrerpolicy": options.Referrerpolicy,
			"target":         options.Target,
			"type":           options.Type,
		},
		InnerHTML:     inner,
		Name:          "a",
		SSLAttributes: map[string][]string{"rel": options.Rel},
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

type QOptions struct {
	Cite string
}

func NewQ(inner string, options QOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{"cite": options.Cite},
		InnerHTML:  inner,
		Name:       "q",
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

type DataOptions struct {
	Value string
}

func NewData(inner string, options DataOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{"value": options.Value},
		InnerHTML:  inner,
		Name:       "data",
	}
}

type TimeOptions struct {
	Datetime string
}

func NewTime(inner string, options TimeOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{"datetime": options.Datetime},
		InnerHTML:  inner,
		Name:       "time",
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

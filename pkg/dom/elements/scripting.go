// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

const (
	Script_Crossorigin_Anonymous      = "anonymous"
	Script_Crossorigin_UseCredentials = "use-credentials"
)

const (
	Script_Referrerpolicy_NoReferrer                  = "no-referrer"
	Script_Referrerpolicy_NoReferrerWhenDowngrade     = "no-referrer-when-downgrade"
	Script_Referrerpolicy_SameOrigin                  = "same-origin"
	Script_Referrerpolicy_Origin                      = "origin"
	Script_Referrerpolicy_StrictOrigin                = "strict-origin"
	Script_Referrerpolicy_OriginWhenCrossOrigin       = "origin-when-cross-origin"
	Script_Referrerpolicy_StrictOriginWhenCrossOrigin = "strict-origin-when-cross-origin"
	Script_Referrerpolicy_UnsafeUrl                   = "unsafe-url"
)

type ScriptOptions struct {
	Src            string
	Type           string
	Nomodule       bool
	Async          bool
	Defer          bool
	Crossorigin    string
	Integrity      string
	Referrerpolicy string
	Blocking       bool
}

func NewScript(inner string, options ScriptOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"crossorigin":    options.Crossorigin,
			"integrity":      options.Integrity,
			"referrerpolicy": options.Referrerpolicy,
			"src":            options.Src,
			"type":           options.Type,
		},
		BoolAttributes: map[string]bool{
			"async":    options.Async,
			"blocking": options.Blocking,
			"defer":    options.Defer,
			"nomodule": options.Nomodule,
		},
		InnerHTML: inner,
		Name:      "script",
	}
}

func NewNoscript(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "noscript",
	}
}

func NewTemplate(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "template",
	}
}

type SlotOptions struct {
	Name string
}

func NewSlot(inner string, options SlotOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{"name": options.Name},
		InnerHTML:  inner,
		Name:       "slot",
	}
}

type CanvasOptions struct {
	Width  int
	Height int
}

func NewCanvas(inner string, options CanvasOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		IntAttributes: map[string]int{
			"height": options.Height,
			"width":  options.Width,
		},
		Name: "canvas",
	}
}

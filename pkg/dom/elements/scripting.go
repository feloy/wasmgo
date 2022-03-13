// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

type ScriptOptions struct {
	Src
	Type
	Nomodule
	Async
	Defer
	Crossorigin
	Integrity
	Referrerpolicy
	Blocking
}

func NewScript(inner string, options ScriptOptions) *dom.Tag {
	return &dom.Tag{
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
	Name
}

func NewSlot(inner string, options SlotOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "slot",
	}
}

type CanvasOptions struct {
	Width
	Height
}

func NewCanvas(inner string, options CanvasOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "canvas",
	}
}

package main

import (
	"syscall/js"

	"github.com/feloy/wasmgo/pkg/dom"
	"github.com/feloy/wasmgo/pkg/dom/elements"
)

type LinkView struct{}

var _ dom.Element = (*LinkView)(nil)

func (o LinkView) Render(document js.Value) js.Value {
	return elements.NewHyperlink("link", elements.HyperlinkOptions{
		Destination: "https://webassembly.org/",
	}).WithId("hyperlink1").Render(document)
}

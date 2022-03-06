package main

import (
	"syscall/js"

	"github.com/feloy/wasmgo/pkg/dom"
)

type LinkView struct{}

var _ dom.Element = (*LinkView)(nil)

func (o LinkView) Render(document js.Value) js.Value {
	return dom.NewHyperlink("link", dom.HyperlinkOptions{
		Destination: "https://webassembly.org/",
	}).WithId("hyperlink1").Render(document)
}

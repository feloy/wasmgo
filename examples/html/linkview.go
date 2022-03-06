package main

import (
	"io"

	"github.com/feloy/wasmgo/pkg/dom"
)

type LinkView struct{}

var _ dom.Element = (*LinkView)(nil)

func (o LinkView) Render(buffer io.Writer) {
	dom.NewHyperlink("link", dom.HyperlinkOptions{
		Destination: "https://webassembly.org/",
	}).WithId("hyperlink1").Render(buffer)
}

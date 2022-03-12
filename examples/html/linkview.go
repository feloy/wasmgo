// +build !js

package main

import (
	"io"

	"github.com/feloy/wasmgo/pkg/dom"
	"github.com/feloy/wasmgo/pkg/dom/elements"
)

type LinkView struct{}

var _ dom.Element = (*LinkView)(nil)

func (o LinkView) Render(buffer io.Writer) {
	elements.NewA("link", elements.AOptions{
		Destination: "https://webassembly.org/",
	}).WithId("hyperlink1").Render(buffer)
}

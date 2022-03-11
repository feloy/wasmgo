package main

import (
	"io"

	"github.com/feloy/wasmgo/pkg/dom"
	"github.com/feloy/wasmgo/pkg/dom/attributes"
	"github.com/feloy/wasmgo/pkg/dom/elements"
)

type PageView struct{}

var _ dom.Element = (*PageView)(nil)

func (o PageView) Render(buffer io.Writer) {
	titleLink := LinkView{}

	elements.NewDiv().
		AddChild(elements.NewHeader1("").WithId("title1").AppendClass("title-level-1").PrependClass("other-class").
			AddChild(elements.NewText("Title 1 (")).
			AddChild(titleLink).
			AddChild(elements.NewText(")"))).
		AddChild(elements.NewHeader2("Title 1.1").WithAutofocus(false).WithTabindex(0)).
		AddChild(elements.NewParagraph("some text: bla bla").WithTabindex(1)).
		AddChild(elements.NewHeader2("Title 1.2")).
		AddChild(elements.NewParagraph("").
			AddChild(elements.NewHyperlink("Web Assembly", elements.HyperlinkOptions{
				Destination: "https://webassembly.org/",
				Relation:    attributes.RelationExternal,
			}))).Render(buffer)
}

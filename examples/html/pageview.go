package main

import (
	"io"

	"github.com/feloy/wasmgo/pkg/dom"
)

type PageView struct{}

var _ dom.Element = (*PageView)(nil)

func (o PageView) Render(buffer io.Writer) {
	titleLink := LinkView{}

	dom.NewDiv().
		AddChild(dom.NewHeader1("").WithId("title1").WithClass("title-level-1").WithClass("other-class").
			AddChild(dom.NewText("Title 1 (")).
			AddChild(titleLink).
			AddChild(dom.NewText(")"))).
		AddChild(dom.NewHeader2("Title 1.1")).
		AddChild(dom.NewParagraph("some text: bla bla")).
		AddChild(dom.NewHeader2("Title 1.2")).
		AddChild(dom.NewParagraph("").
			AddChild(dom.NewHyperlink("Web Assembly", dom.HyperlinkOptions{
				Destination: "https://webassembly.org/",
				Relation:    dom.HyperlinkRelationExternal,
			}))).Render(buffer)
}

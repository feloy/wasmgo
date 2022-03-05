package main

import (
	"github.com/feloy/wasmgo/wasmgo"
	"github.com/feloy/wasmgo/wasmgo/dom"
)

type PageView struct{}

var _ wasmgo.Component = (*PageView)(nil)

func (o PageView) Render() dom.Element {

	titleLink := dom.NewHyperlink("link", dom.HyperlinkOptions{
		Destination: "https://webassembly.org",
	}).WithId("hyperlink1")

	return dom.NewDiv().
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
			})))
}

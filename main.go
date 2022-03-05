package main

import (
	"github.com/feloy/wasmgo/wasmgo"
	"github.com/feloy/wasmgo/wasmgo/dom"
)

func main() {
	pageView := PageView{}
	wasmgo.Render(pageView)
}

type PageView struct{}

func (o PageView) Render() *dom.Tag {
	return dom.NewDiv().
		AddChild(dom.NewHeader1("Title 1: ").AddChild(dom.NewHyperlink("link", dom.HyperlinkOptions{}))).
		AddChild(dom.NewHeader2("Title 1.1")).
		AddChild(dom.NewParagraph("some text: bla bla")).
		AddChild(dom.NewHeader2("Title 1.2")).
		AddChild(dom.NewHyperlink("GNU website", dom.HyperlinkOptions{
			Destination: "http://www.gnu.org",
			Relation:    dom.HyperlinkRelationExternal,
		}))
}

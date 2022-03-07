package main

import (
	"fmt"
	"syscall/js"

	"github.com/feloy/wasmgo/pkg/dom"
	"github.com/feloy/wasmgo/pkg/dom/attributes"
	"github.com/feloy/wasmgo/pkg/dom/elements"
)

type PageView struct{}

var _ dom.Element = (*PageView)(nil)

func (o PageView) Render(document js.Value) js.Value {

	buttonHandler := func() {
		fmt.Println("clicked!")
	}

	titleLink := LinkView{}

	return elements.NewDiv().
		AddChild(elements.NewHeader1("").WithId("title1").WithClass("title-level-1").WithClass("other-class").
			AddChild(elements.NewText("Title 1 (")).
			AddChild(titleLink).
			AddChild(elements.NewText(")"))).
		AddChild(elements.NewHeader2("Title 1.1")).
		AddChild(elements.NewParagraph("some text: bla bla")).
		AddChild(elements.NewButton("click here").AddOnclickHandler(buttonHandler)).
		AddChild(elements.NewHeader2("Title 1.2")).
		AddChild(elements.NewParagraph("").
			AddChild(elements.NewHyperlink("Web Assembly", elements.HyperlinkOptions{
				Destination: "https://webassembly.org/",
				Relation:    attributes.RelationExternal,
			}))).Render(document)
}

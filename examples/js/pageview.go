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

	return elements.NewDiv("").
		AddChild(elements.NewH1("").WithId("title1").AppendClass("title-level-1").AppendClass("other-class").
			AddChild(elements.NewText("Title 1 (")).
			AddChild(titleLink).
			AddChild(elements.NewText(")"))).
		AddChild(elements.NewH2("Title 1.1")).
		AddChild(elements.NewP("some text: bla bla")).
		AddChild(elements.NewButton("click here").AppendAccesskey("k").AddOnclickHandler(buttonHandler)).
		AddChild(elements.NewH2("Title 1.2")).
		AddChild(elements.NewP("").
			AddChild(elements.NewA("Web Assembly", elements.AOptions{
				Destination: "https://webassembly.org/",
				Relation:    attributes.RelationExternal,
			}))).Render(document)
}

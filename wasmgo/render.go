package wasmgo

import (
	"github.com/feloy/wasmgo/wasmgo/dom"
	"github.com/feloy/wasmgo/wasmgo/js"
)

type Component interface {
	Render() dom.Element
}

func Render(jsclient js.Client, element dom.Element) {
	document := jsclient.GetDocument()
	body := jsclient.GetBody(document)
	jsTag := element.Render(document)
	body.Call("appendChild", jsTag)
}

package wasmgo

import (
	"github.com/feloy/wasmgo/wasmgo/dom"
	"github.com/feloy/wasmgo/wasmgo/js"
)

type Component interface {
	Render() dom.Element
}

func Render(jsclient js.Client, component Component) {
	document := jsclient.GetDocument()
	body := jsclient.GetBody(document)
	element := component.Render()
	jsTag := element.Render(document)
	body.Call("appendChild", jsTag)
}

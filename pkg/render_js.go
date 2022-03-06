// +build js

package wasmgo

import (
	"github.com/feloy/wasmgo/pkg/dom"
	"github.com/feloy/wasmgo/pkg/js"
)

func Render(jsclient js.Client, element dom.Element) {
	document := jsclient.GetDocument()
	body := jsclient.GetBody(document)
	jsTag := element.Render(document)
	body.Call("appendChild", jsTag)
}

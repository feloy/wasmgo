package wasmgo

import (
	"syscall/js"

	"github.com/feloy/wasmgo/wasmgo/dom"
	wasmjs "github.com/feloy/wasmgo/wasmgo/js"
)

type Component interface {
	Render() *dom.Tag
}

func Render(component Component) {
	jsclient := wasmjs.NewJS()
	document := jsclient.GetDocument()
	body := jsclient.GetBody(document)
	t := component.Render()
	jsTag := renderTag(document, t)
	body.Call("appendChild", jsTag)

}

func renderTag(document js.Value, tag *dom.Tag) js.Value {
	jsTag := document.Call("createElement", tag.Name)
	jsTag.Set("innerHTML", tag.InnerHTML)
	for attr, value := range tag.Attributes {
		jsTag.Set(attr, value)
	}
	for _, child := range tag.Children {
		jsChild := renderTag(document, &child)
		jsTag.Call("appendChild", jsChild)
	}
	return jsTag
}

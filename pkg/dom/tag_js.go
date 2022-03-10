// +build js

package dom

import (
	"strings"
	"syscall/js"
)

func (tag *Tag) Render(document js.Value) js.Value {
	jsTag := document.Call("createElement", tag.Name)
	jsTag.Set("innerHTML", tag.InnerHTML)
	for attr, value := range tag.Attributes {
		if len(value) > 0 {
			jsTag.Set(attr, value)
		}
	}
	if len(tag.SSLAttributes["class"]) > 0 {
		className := strings.Join(tag.SSLAttributes["class"], " ")
		jsTag.Set("className", className)
	}
	for _, child := range tag.Children {
		jsChild := child.Render(document)
		jsTag.Call("appendChild", jsChild)
	}
	for event, handlers := range tag.Handlers {
		for _, handler := range handlers {
			jsTag.Call("addEventListener", event, js.FuncOf(wrapGoFunction(handler)))
		}
	}
	return jsTag
}

func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface{} {
	return func(_ js.Value, _ []js.Value) interface{} {
		fn()
		return nil
	}
}

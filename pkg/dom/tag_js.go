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
	if len(tag.Classes) > 0 {
		className := strings.Join(tag.Classes, " ")
		jsTag.Set("className", className)
	}
	for _, child := range tag.Children {
		jsChild := child.Render(document)
		jsTag.Call("appendChild", jsChild)
	}
	return jsTag
}

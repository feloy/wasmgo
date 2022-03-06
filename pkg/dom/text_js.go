// +build js

package dom

import (
	"syscall/js"
)

func (o *TextNode) Render(document js.Value) js.Value {
	return document.Call("createTextNode", o.Text)
}

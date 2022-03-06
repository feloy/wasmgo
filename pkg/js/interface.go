// +build js

package js

import "syscall/js"

type Client interface {
	GetBody(document js.Value) js.Value
	GetDocument() js.Value
}

// +build js

package dom

import "syscall/js"

type Element interface {
	Render(document js.Value) js.Value
}

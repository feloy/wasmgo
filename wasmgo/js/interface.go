package js

import "syscall/js"

type Client interface {
	GetBody() js.Value
}

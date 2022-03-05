package js

import (
	"syscall/js"
)

type JS struct{}

func NewJS() *JS {
	return &JS{}
}

func (o JS) GetDocument() js.Value {
	return js.Global().Get("document")
}

func (o JS) GetBody(document js.Value) js.Value {
	body := document.Get("body")
	return body
}

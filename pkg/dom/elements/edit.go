package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewIns() *dom.Tag {
	return &dom.Tag{Name: "ins"}
}
func NewDel() *dom.Tag {
	return &dom.Tag{Name: "del"}
}

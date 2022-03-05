package main

import (
	"github.com/feloy/wasmgo/wasmgo"
	"github.com/feloy/wasmgo/wasmgo/js"
)

func main() {
	pageView := PageView{}
	jsclient := js.NewJS()
	wasmgo.Render(jsclient, pageView)
}

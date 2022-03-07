package main

import (
	wasmgo "github.com/feloy/wasmgo/pkg"
	"github.com/feloy/wasmgo/pkg/js"
)

func main() {
	pageView := PageView{}
	jsclient := js.NewJS()
	wasmgo.Render(jsclient, pageView)
	<-make(chan bool)
}

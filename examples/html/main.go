package main

import (
	"fmt"

	wasmgo "github.com/feloy/wasmgo/pkg"
)

func main() {
	pageView := PageView{}
	fmt.Printf("%s\n", wasmgo.Render(pageView))
}

// +build !js

package wasmgo

import (
	"bytes"

	"github.com/feloy/wasmgo/pkg/dom"
)

func Render(element dom.Element) string {
	var buffer bytes.Buffer
	element.Render(&buffer)
	return buffer.String()
}

// +build !js

package dom

import (
	"io"
)

func (o *TextNode) Render(buffer io.Writer) {
	io.WriteString(buffer, o.Text)
}

// +build !js

package dom

import "io"

type Element interface {
	Render(buffer io.Writer)
}

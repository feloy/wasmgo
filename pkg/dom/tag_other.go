// +build !js

package dom

import (
	"fmt"
	"html"
	"io"
	"strings"
)

func (tag *Tag) Render(buffer io.Writer) {
	fmt.Fprintf(buffer, "<%s", tag.Name)
	for attr, value := range tag.Attributes {
		if len(value) > 0 {
			fmt.Fprintf(buffer, ` %s="%s"`, attr, html.EscapeString(value))
		}
	}
	if len(tag.Classes) > 0 {
		className := strings.Join(tag.Classes, " ")
		fmt.Fprintf(buffer, ` class="%s"`, className)
	}
	fmt.Fprintf(buffer, ">%s", html.EscapeString(tag.InnerHTML))
	for _, child := range tag.Children {
		child.Render(buffer)
	}
	fmt.Fprintf(buffer, "</%s>", tag.Name)
}

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
	for attr, value := range tag.BoolAttributes {
		if value {
			fmt.Fprintf(buffer, ` %s`, attr)
		}
	}
	for attr, value := range tag.IntAttributes {
		fmt.Fprintf(buffer, ` %s="%d"`, attr, value)
	}
	for attr, values := range tag.SSLAttributes {
		list := strings.Join(values, " ")
		fmt.Fprintf(buffer, ` %s="%s"`, attr, list)
	}
	fmt.Fprintf(buffer, ">%s", html.EscapeString(tag.InnerHTML))
	for _, child := range tag.Children {
		child.Render(buffer)
	}
	fmt.Fprintf(buffer, "</%s>", tag.Name)
}

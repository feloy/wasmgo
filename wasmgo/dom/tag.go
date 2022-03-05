package dom

import (
	"strings"
	"syscall/js"
)

type Tag struct {
	Name       string
	InnerHTML  string
	Attributes map[string]string
	Classes    []string
	Children   []Element
}

func (o *Tag) appendAttribute(key, value string) {
	if o.Attributes == nil {
		o.Attributes = make(map[string]string)
	}
	o.Attributes[key] = value
}

// WithId sets the "id" attribute of the tag
func (o *Tag) WithId(id string) *Tag {
	o.appendAttribute("id", id)
	return o
}

// WithClass adds a class to the tag
func (o *Tag) WithClass(class string) *Tag {
	o.Classes = append(o.Classes, class)
	return o
}

// AddChild adds a child tag to the tag
func (o *Tag) AddChild(child Element) *Tag {
	o.Children = append(o.Children, child)
	return o
}

func (tag *Tag) Render(document js.Value) js.Value {
	jsTag := document.Call("createElement", tag.Name)
	jsTag.Set("innerHTML", tag.InnerHTML)
	for attr, value := range tag.Attributes {
		jsTag.Set(attr, value)
	}
	if len(tag.Classes) > 0 {
		className := strings.Join(tag.Classes, " ")
		jsTag.Set("className", className)
	}
	for _, child := range tag.Children {
		jsChild := child.Render(document)
		jsTag.Call("appendChild", jsChild)
	}
	return jsTag
}

// NewDiv returns a new empty "<div>" tag
func NewDiv() *Tag {
	return &Tag{
		Name: "div",
	}
}

// NewHeader1 returns a new "<h1>" tag with title content
func NewHeader1(title string) *Tag {
	return &Tag{
		Name:      "h1",
		InnerHTML: title,
	}
}

// NewHeader2 returns a new "<h2>" tag with title content
func NewHeader2(title string) *Tag {
	return &Tag{
		Name:      "h2",
		InnerHTML: title,
	}
}

// NewParagraph returns a new "<p>" tag with text content
func NewParagraph(text string) *Tag {
	return &Tag{
		Name:      "p",
		InnerHTML: text,
	}
}

package dom

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

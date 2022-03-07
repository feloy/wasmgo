package dom

type EventHandler func()

type Tag struct {
	Name       string
	InnerHTML  string
	Attributes map[string]string
	Classes    []string
	Children   []Element
	Handlers   map[string][]EventHandler
}

func (o *Tag) appendAttribute(key, value string) {
	if o.Attributes == nil {
		o.Attributes = make(map[string]string)
	}
	o.Attributes[key] = value
}

func (o *Tag) addHandler(name string, handler EventHandler) {
	if o.Handlers == nil {
		o.Handlers = make(map[string][]EventHandler)
	}
	o.Handlers[name] = append(o.Handlers[name], handler)
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

func (o *Tag) AddOnclickHandler(handler EventHandler) *Tag {
	o.addHandler("click", handler)
	return o
}

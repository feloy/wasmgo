package dom

type EventHandler func()

type Tag struct {
	Name            string
	InnerHTML       string
	Attributes      map[string]string
	BoolAttributes  map[string]bool
	IntAttributes   map[string]int
	FloatAttributes map[string]float32
	SSLAttributes   map[string][]string
	Children        []Element
	Handlers        map[string][]EventHandler
	OmitEndTag      bool
}

func (o *Tag) appendAttribute(key, value string) {
	if o.Attributes == nil {
		o.Attributes = make(map[string]string)
	}
	o.Attributes[key] = value
}

func (o *Tag) appendBoolAttribute(key string, value bool) {
	if o.BoolAttributes == nil {
		o.BoolAttributes = make(map[string]bool)
	}
	o.BoolAttributes[key] = value
}

func (o *Tag) appendIntAttribute(key string, value int) {
	if o.IntAttributes == nil {
		o.IntAttributes = make(map[string]int)
	}
	o.IntAttributes[key] = value
}

func (o *Tag) appendSSLAttribute(key, value string) {
	if o.SSLAttributes == nil {
		o.SSLAttributes = make(map[string][]string)
	}
	o.SSLAttributes[key] = append(o.SSLAttributes[key], value)
}

func (o *Tag) prependSSLAttribute(key, value string) {
	if o.SSLAttributes == nil {
		o.SSLAttributes = make(map[string][]string)
	}
	o.SSLAttributes[key] = append([]string{value}, o.SSLAttributes[key]...)
}

func (o *Tag) addHandler(name string, handler EventHandler) {
	if o.Handlers == nil {
		o.Handlers = make(map[string][]EventHandler)
	}
	o.Handlers[name] = append(o.Handlers[name], handler)
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

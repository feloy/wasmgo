package dom

func (o *Tag) AppendClass(class string) *Tag {
	o.appendSSLAttribute("class", class)
	return o
}
func (o *Tag) PrependClass(class string) *Tag {
	o.prependSSLAttribute("class", class)
	return o
}

func (o *Tag) WithId(id string) *Tag {
	o.appendAttribute("id", id)
	return o
}

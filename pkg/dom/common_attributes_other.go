// +build !js

package dom

func (o *Tag) AppendAccesskey(accesskey string) *Tag {
	o.appendSSLAttribute("accesskey", accesskey)
	return o
}
func (o *Tag) PrependAccesskey(accesskey string) *Tag {
	o.prependSSLAttribute("accesskey", accesskey)
	return o
}

func (o *Tag) AppendClass(class string) *Tag {
	o.appendSSLAttribute("class", class)
	return o
}
func (o *Tag) PrependClass(class string) *Tag {
	o.prependSSLAttribute("class", class)
	return o
}

func (o *Tag) WithContenteditable(contenteditable string) *Tag {
	o.appendAttribute("contenteditable", contenteditable)
	return o
}

func (o *Tag) WithEnterkeyhint(enterkeyhint string) *Tag {
	o.appendAttribute("enterkeyhint", enterkeyhint)
	return o
}

func (o *Tag) WithInputmode(inputmode string) *Tag {
	o.appendAttribute("inputmode", inputmode)
	return o
}

func (o *Tag) WithTabindex(tabindex int) *Tag {
	o.appendIntAttribute("tabindex", tabindex)
	return o
}

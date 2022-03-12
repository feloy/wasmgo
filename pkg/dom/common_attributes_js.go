// +build js
// This file is generated. Please do not edit

package dom

func (o *Tag) AppendAccesskey(accesskey string) *Tag {
	o.appendSSLAttribute("accessKey", accesskey)
	return o
}
func (o *Tag) PrependAccesskey(accesskey string) *Tag {
	o.prependSSLAttribute("accessKey", accesskey)
	return o
}

func (o *Tag) AppendClass(class string) *Tag {
	o.appendSSLAttribute("className", class)
	return o
}
func (o *Tag) PrependClass(class string) *Tag {
	o.prependSSLAttribute("className", class)
	return o
}

func (o *Tag) WithContenteditable(contenteditable string) *Tag {
	o.appendAttribute("contentEditable", contenteditable)
	return o
}

func (o *Tag) WithEnterkeyhint(enterkeyhint string) *Tag {
	o.appendAttribute("enterKeyHint", enterkeyhint)
	return o
}

func (o *Tag) WithInputmode(inputmode string) *Tag {
	o.appendAttribute("inputMode", inputmode)
	return o
}

func (o *Tag) WithTabindex(tabindex int) *Tag {
	o.appendIntAttribute("tabIndex", tabindex)
	return o
}

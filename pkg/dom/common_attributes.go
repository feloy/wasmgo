package dom

func (o *Tag) WithAutocapitalize(autocapitalize string) *Tag {
	o.appendAttribute("autocapitalize", autocapitalize)
	return o
}

func (o *Tag) WithAutofocus(autofocus bool) *Tag {
	o.appendBoolAttribute("autofocus", autofocus)
	return o
}

func (o *Tag) WithDir(dir string) *Tag {
	o.appendAttribute("dir", dir)
	return o
}

func (o *Tag) WithDraggable(draggable string) *Tag {
	o.appendAttribute("draggable", draggable)
	return o
}

func (o *Tag) WithHidden(hidden bool) *Tag {
	o.appendBoolAttribute("hidden", hidden)
	return o
}

func (o *Tag) WithId(id string) *Tag {
	o.appendAttribute("id", id)
	return o
}

func (o *Tag) WithIs(is string) *Tag {
	o.appendAttribute("is", is)
	return o
}

func (o *Tag) WithItemid(itemid string) *Tag {
	o.appendAttribute("itemid", itemid)
	return o
}

func (o *Tag) AppendItemprop(itemprop string) *Tag {
	o.appendSSLAttribute("itemprop", itemprop)
	return o
}
func (o *Tag) PrependItemprop(itemprop string) *Tag {
	o.prependSSLAttribute("itemprop", itemprop)
	return o
}

func (o *Tag) AppendItemref(itemref string) *Tag {
	o.appendSSLAttribute("itemref", itemref)
	return o
}
func (o *Tag) PrependItemref(itemref string) *Tag {
	o.prependSSLAttribute("itemref", itemref)
	return o
}

func (o *Tag) WithItemscope(itemscope bool) *Tag {
	o.appendBoolAttribute("itemscope", itemscope)
	return o
}

func (o *Tag) AppendItemtype(itemtype string) *Tag {
	o.appendSSLAttribute("itemtype", itemtype)
	return o
}
func (o *Tag) PrependItemtype(itemtype string) *Tag {
	o.prependSSLAttribute("itemtype", itemtype)
	return o
}

func (o *Tag) WithLang(lang string) *Tag {
	o.appendAttribute("lang", lang)
	return o
}

func (o *Tag) WithNonce(nonce string) *Tag {
	o.appendAttribute("nonce", nonce)
	return o
}

func (o *Tag) WithRole(role string) *Tag {
	o.appendAttribute("role", role)
	return o
}

func (o *Tag) WithSlot(slot string) *Tag {
	o.appendAttribute("slot", slot)
	return o
}

func (o *Tag) WithSpellcheck(spellcheck string) *Tag {
	o.appendAttribute("spellcheck", spellcheck)
	return o
}

func (o *Tag) WithStyle(style string) *Tag {
	o.appendAttribute("style", style)
	return o
}

func (o *Tag) WithTitle(title string) *Tag {
	o.appendAttribute("title", title)
	return o
}

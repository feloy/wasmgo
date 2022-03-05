package dom

type Tag struct {
	Name       string
	InnerHTML  string
	Attributes map[string]string
	Children   []Tag
}

func NewDiv() *Tag {
	return &Tag{
		Name: "div",
	}
}

func NewHeader1(title string) *Tag {
	return &Tag{
		Name:      "h1",
		InnerHTML: title,
	}
}

func NewHeader2(title string) *Tag {
	return &Tag{
		Name:      "h2",
		InnerHTML: title,
	}
}

func NewParagraph(text string) *Tag {
	return &Tag{
		Name:      "p",
		InnerHTML: text,
	}
}

func (o *Tag) AddChild(child *Tag) *Tag {
	o.Children = append(o.Children, *child)
	return o
}

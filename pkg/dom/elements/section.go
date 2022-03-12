package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewBody() *dom.Tag {
	return &dom.Tag{Name: "body"}
}
func NewArticle() *dom.Tag {
	return &dom.Tag{Name: "article"}
}
func NewSection() *dom.Tag {
	return &dom.Tag{Name: "section"}
}
func NewNav() *dom.Tag {
	return &dom.Tag{Name: "nav"}
}
func NewAside() *dom.Tag {
	return &dom.Tag{Name: "aside"}
}
func NewH1(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h1",
	}
}
func NewH2(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h2",
	}
}
func NewH3(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h3",
	}
}
func NewH4(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h4",
	}
}
func NewH5(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h5",
	}
}
func NewH6(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "h6",
	}
}
func NewHgroup() *dom.Tag {
	return &dom.Tag{Name: "hgroup"}
}
func NewHeader() *dom.Tag {
	return &dom.Tag{Name: "header"}
}
func NewFooter() *dom.Tag {
	return &dom.Tag{Name: "footer"}
}
func NewAddress() *dom.Tag {
	return &dom.Tag{Name: "address"}
}

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewPicture() *dom.Tag {
	return &dom.Tag{Name: "picture"}
}
func NewSource() *dom.Tag {
	return &dom.Tag{Name: "source"}
}
func NewImg() *dom.Tag {
	return &dom.Tag{Name: "img"}
}
func NewIframe() *dom.Tag {
	return &dom.Tag{Name: "iframe"}
}
func NewEmbed() *dom.Tag {
	return &dom.Tag{Name: "embed"}
}
func NewObject() *dom.Tag {
	return &dom.Tag{Name: "object"}
}
func NewParam() *dom.Tag {
	return &dom.Tag{Name: "param"}
}
func NewVideo() *dom.Tag {
	return &dom.Tag{Name: "video"}
}
func NewAudio() *dom.Tag {
	return &dom.Tag{Name: "audio"}
}
func NewTrack() *dom.Tag {
	return &dom.Tag{Name: "track"}
}
func NewMap() *dom.Tag {
	return &dom.Tag{Name: "map"}
}
func NewArea() *dom.Tag {
	return &dom.Tag{Name: "area"}
}

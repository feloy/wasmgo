// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewPicture() *dom.Tag {
	return &dom.Tag{Name: "picture"}
}
func NewSource() *dom.Tag {
	return &dom.Tag{
		Name:       "source",
		OmitEndTag: true,
	}
}
func NewImg() *dom.Tag {
	return &dom.Tag{
		Name:       "img",
		OmitEndTag: true,
	}
}
func NewIframe() *dom.Tag {
	return &dom.Tag{Name: "iframe"}
}
func NewEmbed() *dom.Tag {
	return &dom.Tag{
		Name:       "embed",
		OmitEndTag: true,
	}
}
func NewObject(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "object",
	}
}
func NewParam() *dom.Tag {
	return &dom.Tag{
		Name:       "param",
		OmitEndTag: true,
	}
}
func NewVideo(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "video",
	}
}
func NewAudio(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "audio",
	}
}
func NewTrack() *dom.Tag {
	return &dom.Tag{
		Name:       "track",
		OmitEndTag: true,
	}
}
func NewMap(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "map",
	}
}
func NewArea() *dom.Tag {
	return &dom.Tag{
		Name:       "area",
		OmitEndTag: true,
	}
}

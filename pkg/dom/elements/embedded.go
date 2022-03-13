// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewPicture() *dom.Tag {
	return &dom.Tag{Name: "picture"}
}

type SourceOptions struct {
	Type
	Src
	Srcset
	Sizes
	Media
	Width
	Height
}

func NewSource(options SourceOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "source",
		OmitEndTag: true,
	}
}

type ImgOptions struct {
	Alt
	Src
	Srcset
	Sizes
	Crossorigin
	Usemap
	Ismap
	Width
	Height
	Referrerpolicy
	Decoding
	Loading
}

func NewImg(options ImgOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "img",
		OmitEndTag: true,
	}
}

type IframeOptions struct {
	Src
	Srcset
	Name
	Sandbox
	Allow
	Allowfullscreen
	Width
	Height
	Referrerpolicy
	Loading
}

func NewIframe(options IframeOptions) *dom.Tag {
	return &dom.Tag{Name: "iframe"}
}

type EmbedOptions struct {
	Src
	Type
	Width
	Height
}

func NewEmbed(options EmbedOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "embed",
		OmitEndTag: true,
	}
}

type ObjectOptions struct {
	Data
	Type
	Name
	Form
	Width
	Height
}

func NewObject(inner string, options ObjectOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "object",
	}
}

type ParamOptions struct {
	Name
	Value
}

func NewParam(options ParamOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "param",
		OmitEndTag: true,
	}
}

type VideoOptions struct {
	Src
	Crossorigin
	Poster
	Preload
	Autoplay
	Playsinline
	Loop
	Muted
	Controls
	Width
	Height
}

func NewVideo(inner string, options VideoOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "video",
	}
}

type AudioOptions struct {
	Src
	Crossorigin
	Preload
	Autoplay
	Loop
	Muted
	Controls
}

func NewAudio(inner string, options AudioOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "audio",
	}
}

type TrackOptions struct {
	Kind
	Src
	Srclang
	Label
	Default
}

func NewTrack(options TrackOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "track",
		OmitEndTag: true,
	}
}

type MapOptions struct {
	Name
}

func NewMap(inner string, options MapOptions) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "map",
	}
}

type AreaOptions struct {
	Alt
	Coords
	Shape
	Href
	Target
	Download
	Ping
	Rel
	Referrerpolicy
}

func NewArea(options AreaOptions) *dom.Tag {
	return &dom.Tag{
		Name:       "area",
		OmitEndTag: true,
	}
}

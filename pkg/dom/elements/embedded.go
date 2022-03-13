// This file is generated. Please do not edit

package elements

import "github.com/feloy/wasmgo/pkg/dom"

func NewPicture() *dom.Tag {
	return &dom.Tag{Name: "picture"}
}

type SourceOptions struct {
	Type   string
	Src    string
	Srcset string
	Sizes  string
	Media  string
	Width  int
	Height int
}

func NewSource(options SourceOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"media":  options.Media,
			"sizes":  options.Sizes,
			"src":    options.Src,
			"srcset": options.Srcset,
			"type":   options.Type,
		},
		IntAttributes: map[string]int{
			"height": options.Height,
			"width":  options.Width,
		},
		Name:       "source",
		OmitEndTag: true,
	}
}

const (
	Img_Crossorigin_Anonymous      = "anonymous"
	Img_Crossorigin_UseCredentials = "use-credentials"
)

const (
	Img_Referrerpolicy_NoReferrer                  = "no-referrer"
	Img_Referrerpolicy_NoReferrerWhenDowngrade     = "no-referrer-when-downgrade"
	Img_Referrerpolicy_SameOrigin                  = "same-origin"
	Img_Referrerpolicy_Origin                      = "origin"
	Img_Referrerpolicy_StrictOrigin                = "strict-origin"
	Img_Referrerpolicy_OriginWhenCrossOrigin       = "origin-when-cross-origin"
	Img_Referrerpolicy_StrictOriginWhenCrossOrigin = "strict-origin-when-cross-origin"
	Img_Referrerpolicy_UnsafeUrl                   = "unsafe-url"
)

const (
	Img_Decoding_Sync  = "sync"
	Img_Decoding_Async = "async"
	Img_Decoding_Auto  = "auto"
)

const (
	Img_Loading_Lazy  = "lazy"
	Img_Loading_Eager = "eager"
)

type ImgOptions struct {
	Alt            string
	Src            string
	Srcset         string
	Sizes          string
	Crossorigin    string
	Usemap         string
	Ismap          bool
	Width          int
	Height         int
	Referrerpolicy string
	Decoding       string
	Loading        string
}

func NewImg(options ImgOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"alt":            options.Alt,
			"crossorigin":    options.Crossorigin,
			"decoding":       options.Decoding,
			"loading":        options.Loading,
			"referrerpolicy": options.Referrerpolicy,
			"sizes":          options.Sizes,
			"src":            options.Src,
			"srcset":         options.Srcset,
			"usemap":         options.Usemap,
		},
		BoolAttributes: map[string]bool{"ismap": options.Ismap},
		IntAttributes: map[string]int{
			"height": options.Height,
			"width":  options.Width,
		},
		Name:       "img",
		OmitEndTag: true,
	}
}

const (
	Iframe_Sandbox_AllowDownloads                      = "allow-downloads"
	Iframe_Sandbox_AllowForms                          = "allow-forms"
	Iframe_Sandbox_AllowModals                         = "allow-modals"
	Iframe_Sandbox_AllowOrientationLock                = "allow-orientation-lock"
	Iframe_Sandbox_AllowPointerLock                    = "allow-pointer-lock"
	Iframe_Sandbox_AllowPopups                         = "allow-popups"
	Iframe_Sandbox_AllowPopupsToEscapeSandbox          = "allow-popups-to-escape-sandbox"
	Iframe_Sandbox_AllowPresentation                   = "allow-presentation"
	Iframe_Sandbox_AllowSameOrigin                     = "allow-same-origin"
	Iframe_Sandbox_AllowScripts                        = "allow-scripts"
	Iframe_Sandbox_AllowTopNavigation                  = "allow-top-navigation"
	Iframe_Sandbox_AllowTopNavigationByUserActivation  = "allow-top-navigation-by-user-activation"
	Iframe_Sandbox_AllowTopNavigationToCustomProtocols = "allow-top-navigation-to-custom-protocols"
)

const (
	Iframe_Referrerpolicy_NoReferrer                  = "no-referrer"
	Iframe_Referrerpolicy_NoReferrerWhenDowngrade     = "no-referrer-when-downgrade"
	Iframe_Referrerpolicy_SameOrigin                  = "same-origin"
	Iframe_Referrerpolicy_Origin                      = "origin"
	Iframe_Referrerpolicy_StrictOrigin                = "strict-origin"
	Iframe_Referrerpolicy_OriginWhenCrossOrigin       = "origin-when-cross-origin"
	Iframe_Referrerpolicy_StrictOriginWhenCrossOrigin = "strict-origin-when-cross-origin"
	Iframe_Referrerpolicy_UnsafeUrl                   = "unsafe-url"
)

const (
	Iframe_Loading_Lazy  = "lazy"
	Iframe_Loading_Eager = "eager"
)

type IframeOptions struct {
	Src             string
	Srcset          string
	Name            string
	Sandbox         []string
	Allow           string
	Allowfullscreen bool
	Width           int
	Height          int
	Referrerpolicy  string
	Loading         string
}

func NewIframe(options IframeOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"allow":          options.Allow,
			"loading":        options.Loading,
			"name":           options.Name,
			"referrerpolicy": options.Referrerpolicy,
			"src":            options.Src,
			"srcset":         options.Srcset,
		},
		BoolAttributes: map[string]bool{"allowfullscreen": options.Allowfullscreen},
		IntAttributes: map[string]int{
			"height": options.Height,
			"width":  options.Width,
		},
		Name:          "iframe",
		SSLAttributes: map[string][]string{"sandbox": options.Sandbox},
	}
}

type EmbedOptions struct {
	Src    string
	Type   string
	Width  int
	Height int
}

func NewEmbed(options EmbedOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"src":  options.Src,
			"type": options.Type,
		},
		IntAttributes: map[string]int{
			"height": options.Height,
			"width":  options.Width,
		},
		Name:       "embed",
		OmitEndTag: true,
	}
}

type ObjectOptions struct {
	Data   string
	Type   string
	Name   string
	Form   string
	Width  int
	Height int
}

func NewObject(inner string, options ObjectOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"data": options.Data,
			"form": options.Form,
			"name": options.Name,
			"type": options.Type,
		},
		InnerHTML: inner,
		IntAttributes: map[string]int{
			"height": options.Height,
			"width":  options.Width,
		},
		Name: "object",
	}
}

type ParamOptions struct {
	Name  string
	Value string
}

func NewParam(options ParamOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"name":  options.Name,
			"value": options.Value,
		},
		Name:       "param",
		OmitEndTag: true,
	}
}

const (
	Video_Crossorigin_Anonymous      = "anonymous"
	Video_Crossorigin_UseCredentials = "use-credentials"
)

const (
	Video_Preload_None     = "none"
	Video_Preload_Metadata = "metadata"
	Video_Preload_Auto     = "auto"
)

type VideoOptions struct {
	Src         string
	Crossorigin string
	Poster      string
	Preload     string
	Autoplay    bool
	Playsinline bool
	Loop        bool
	Muted       bool
	Controls    bool
	Width       int
	Height      int
}

func NewVideo(inner string, options VideoOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"crossorigin": options.Crossorigin,
			"poster":      options.Poster,
			"preload":     options.Preload,
			"src":         options.Src,
		},
		BoolAttributes: map[string]bool{
			"autoplay":    options.Autoplay,
			"controls":    options.Controls,
			"loop":        options.Loop,
			"muted":       options.Muted,
			"playsinline": options.Playsinline,
		},
		InnerHTML: inner,
		IntAttributes: map[string]int{
			"height": options.Height,
			"width":  options.Width,
		},
		Name: "video",
	}
}

const (
	Audio_Crossorigin_Anonymous      = "anonymous"
	Audio_Crossorigin_UseCredentials = "use-credentials"
)

const (
	Audio_Preload_None     = "none"
	Audio_Preload_Metadata = "metadata"
	Audio_Preload_Auto     = "auto"
)

type AudioOptions struct {
	Src         string
	Crossorigin string
	Preload     string
	Autoplay    bool
	Loop        bool
	Muted       bool
	Controls    bool
}

func NewAudio(inner string, options AudioOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"crossorigin": options.Crossorigin,
			"preload":     options.Preload,
			"src":         options.Src,
		},
		BoolAttributes: map[string]bool{
			"autoplay": options.Autoplay,
			"controls": options.Controls,
			"loop":     options.Loop,
			"muted":    options.Muted,
		},
		InnerHTML: inner,
		Name:      "audio",
	}
}

const (
	Track_Kind_Subtitles    = "subtitles"
	Track_Kind_Captions     = "captions"
	Track_Kind_Descriptions = "descriptions"
	Track_Kind_Chapters     = "chapters"
	Track_Kind_Metadata     = "metadata"
)

type TrackOptions struct {
	Kind    string
	Src     string
	Srclang string
	Label   string
	Default bool
}

func NewTrack(options TrackOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"kind":    options.Kind,
			"label":   options.Label,
			"src":     options.Src,
			"srclang": options.Srclang,
		},
		BoolAttributes: map[string]bool{"default": options.Default},
		Name:           "track",
		OmitEndTag:     true,
	}
}

type MapOptions struct {
	Name string
}

func NewMap(inner string, options MapOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{"name": options.Name},
		InnerHTML:  inner,
		Name:       "map",
	}
}

const (
	Area_Shape_Circle  = "circle"
	Area_Shape_Default = "default"
	Area_Shape_Poly    = "poly"
	Area_Shape_Rect    = "rect"
)

const (
	Area_Rel_Noopener   = "noopener"
	Area_Rel_Bookmark   = "bookmark"
	Area_Rel_Nofollow   = "nofollow"
	Area_Rel_Author     = "author"
	Area_Rel_Opener     = "opener"
	Area_Rel_Tag        = "tag"
	Area_Rel_Prev       = "prev"
	Area_Rel_Alternate  = "alternate"
	Area_Rel_External   = "external"
	Area_Rel_Help       = "help"
	Area_Rel_License    = "license"
	Area_Rel_Next       = "next"
	Area_Rel_Noreferrer = "noreferrer"
	Area_Rel_Search     = "search"
)

const (
	Area_Referrerpolicy_NoReferrer                  = "no-referrer"
	Area_Referrerpolicy_NoReferrerWhenDowngrade     = "no-referrer-when-downgrade"
	Area_Referrerpolicy_SameOrigin                  = "same-origin"
	Area_Referrerpolicy_Origin                      = "origin"
	Area_Referrerpolicy_StrictOrigin                = "strict-origin"
	Area_Referrerpolicy_OriginWhenCrossOrigin       = "origin-when-cross-origin"
	Area_Referrerpolicy_StrictOriginWhenCrossOrigin = "strict-origin-when-cross-origin"
	Area_Referrerpolicy_UnsafeUrl                   = "unsafe-url"
)

type AreaOptions struct {
	Alt            string
	Coords         string
	Shape          string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            []string
	Referrerpolicy string
}

func NewArea(options AreaOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"alt":            options.Alt,
			"coords":         options.Coords,
			"download":       options.Download,
			"href":           options.Href,
			"ping":           options.Ping,
			"referrerpolicy": options.Referrerpolicy,
			"shape":          options.Shape,
			"target":         options.Target,
		},
		Name:          "area",
		OmitEndTag:    true,
		SSLAttributes: map[string][]string{"rel": options.Rel},
	}
}

// +build !js

package main

import (
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

type category struct {
	name     string
	elements []element
}

type element struct {
	name       string
	hasInner   bool
	noEndTag   bool
	attributes []attribute
}

type attribute struct {
	name     string
	property string
	key      string
	typ      string
	values   []string
}

var categories = []category{
	{
		name: "metadata",
		elements: []element{
			{
				name:     "head",
				hasInner: false,
			},
			{
				name:     "title",
				hasInner: true,
			},
			{
				name:     "base",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "link",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "meta",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "style",
				hasInner: true,
			},
		},
	},
	{
		name: "section",
		elements: []element{
			{
				name:     "body",
				hasInner: true,
			},
			{
				name:     "article",
				hasInner: true,
			},
			{
				name:     "section",
				hasInner: true,
			},
			{
				name:     "nav",
				hasInner: true,
			},
			{
				name:     "aside",
				hasInner: true,
			},
			{
				name:     "h1",
				hasInner: true,
			},
			{
				name:     "h2",
				hasInner: true,
			},
			{
				name:     "h3",
				hasInner: true,
			},
			{
				name:     "h4",
				hasInner: true,
			},
			{
				name:     "h5",
				hasInner: true,
			},
			{
				name:     "h6",
				hasInner: true,
			},
			{
				name:     "hgroup",
				hasInner: true,
			},
			{
				name:     "header",
				hasInner: true,
			},
			{
				name:     "footer",
				hasInner: true,
			},
			{
				name:     "address",
				hasInner: true,
			},
		},
	},
	{
		name: "grouping",
		elements: []element{
			{
				name:     "p",
				hasInner: true,
			},
			{
				name:     "hr",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "pre",
				hasInner: true,
			},
			{
				name:     "blockquote",
				hasInner: true,
				attributes: []attribute{
					{
						name: "cite",
						key:  "cite",
						typ:  "string",
					},
				},
			},
			{
				name:     "ol",
				hasInner: false,
				attributes: []attribute{
					{
						name: "reversed",
						key:  "reversed",
						typ:  BOOL,
					},
					{
						name: "start",
						key:  "start",
						typ:  INTEGER,
					},
					{
						name:   "type",
						key:    "type",
						typ:    STRING,
						values: []string{"1", "a", "A", "i", "I"},
					},
				},
			},
			{
				name:     "ul",
				hasInner: false,
			},
			{
				name:     "menu",
				hasInner: false,
			},
			{
				name:     "li",
				hasInner: true,
				attributes: []attribute{
					{
						name: "value",
						key:  "value",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "dl",
				hasInner: false,
			},
			{
				name:     "dt",
				hasInner: true,
			},
			{
				name:     "dd",
				hasInner: true,
			},
			{
				name:     "figure",
				hasInner: true,
			},
			{
				name:     "figcaption",
				hasInner: true,
			},
			{
				name:     "main",
				hasInner: true,
			},
			{
				name:     "div",
				hasInner: true,
			},
		},
	},
	{
		name: "text",
		elements: []element{
			{
				name:     "a",
				hasInner: true,
				attributes: []attribute{
					{
						name: "href",
						key:  "href",
						typ:  STRING,
					},
					{
						name: "target",
						key:  "target",
						typ:  STRING,
					},
					{
						name: "download",
						key:  "download",
						typ:  STRING,
					},
					{
						name: "ping",
						key:  "ping",
						typ:  STRING,
					},
					{
						name:   "rel",
						key:    "rel",
						typ:    SPACE_SEPARATED_LIST,
						values: relAttributes(REL_A),
					},
					{
						name: "hreflang",
						key:  "hreflang",
						typ:  STRING, /* TODO: BCP 47 */
					},
					{
						name: "type",
						key:  "type",
						typ:  STRING,
					},
					{
						name:   "referrerpolicy",
						key:    "referrerpolicy",
						typ:    STRING,
						values: referrerPolicyValues,
					},
				},
			},
			{
				name:     "em",
				hasInner: true,
			},
			{
				name:     "strong",
				hasInner: true,
			},
			{
				name:     "small",
				hasInner: true,
			},
			{
				name:     "s",
				hasInner: true,
			},
			{
				name:     "cite",
				hasInner: true,
			},
			{
				name:     "q",
				hasInner: true,
				attributes: []attribute{
					{
						name: "cite",
						key:  "cite",
						typ:  STRING,
					},
				},
			},
			{
				name:     "dfn",
				hasInner: true,
			},
			{
				name:     "abbr",
				hasInner: true,
			},
			{
				name:     "ruby",
				hasInner: true,
			},
			{
				name:     "rt",
				hasInner: true,
			},
			{
				name:     "rp",
				hasInner: true,
			},
			{
				name:     "data",
				hasInner: true,
				attributes: []attribute{
					{
						name: "value",
						key:  "value",
						typ:  STRING,
					},
				},
			},
			{
				name:     "time",
				hasInner: true,
				attributes: []attribute{
					{
						name: "datetime",
						key:  "datetime",
						typ:  STRING,
					},
				},
			},
			{
				name:     "code",
				hasInner: true,
			},
			{
				name:     "var",
				hasInner: true,
			},
			{
				name:     "samp",
				hasInner: true,
			},
			{
				name:     "kbd",
				hasInner: true,
			},
			{
				name:     "sub",
				hasInner: true,
			},
			{
				name:     "sup",
				hasInner: true,
			},
			{
				name:     "i",
				hasInner: true,
			},
			{
				name:     "b",
				hasInner: true,
			},
			{
				name:     "u",
				hasInner: true,
			},
			{
				name:     "mark",
				hasInner: true,
			},
			{
				name:     "bdi",
				hasInner: true,
			},
			{
				name:     "bdo",
				hasInner: true,
			},
			{
				name:     "span",
				hasInner: true,
			},
			{
				name:     "br",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "wbr",
				hasInner: false,
				noEndTag: true,
			},
		},
	},
	{
		name: "edit",
		elements: []element{
			{
				name:     "ins",
				hasInner: true,
				attributes: []attribute{
					{
						name: "cite",
						key:  "cite",
						typ:  STRING,
					},
					{
						name: "datetime",
						key:  "datetime",
						typ:  STRING,
					},
				},
			},
			{
				name:     "del",
				hasInner: true,
				attributes: []attribute{
					{
						name: "cite",
						key:  "cite",
						typ:  STRING,
					},
					{
						name: "datetime",
						key:  "datetime",
						typ:  STRING,
					},
				},
			},
		},
	},
	{
		name: "embedded",
		elements: []element{
			{
				name:     "picture",
				hasInner: false,
			},
			{
				name:     "source",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "type",
						key:  "type",
						typ:  STRING,
					},
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "srcset",
						key:  "srcset",
						typ:  STRING,
					},
					{
						name: "sizes",
						key:  "sizes",
						typ:  STRING,
					},
					{
						name: "media",
						key:  "media",
						typ:  STRING,
					},
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "img",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "alt",
						key:  "alt",
						typ:  STRING,
					},
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "srcset",
						key:  "srcset",
						typ:  STRING,
					},
					{
						name: "sizes",
						key:  "sizes",
						typ:  STRING,
					},
					{
						name: "crossorigin",
						key:  "crossorigin",
						typ:  STRING,
						values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						name: "usemap",
						key:  "usemap",
						typ:  STRING,
					},
					{
						name: "ismap",
						key:  "ismap",
						typ:  BOOL,
					},
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
					{
						name:   "referrerpolicy",
						key:    "referrerpolicy",
						typ:    STRING,
						values: referrerPolicyValues,
					},
					{
						name:   "decoding",
						key:    "decoding",
						typ:    STRING,
						values: []string{"sync", "async", "auto"},
					},
					{
						name:   "loading",
						key:    "loading",
						typ:    STRING,
						values: []string{"lazy", "eager"},
					},
				},
			},
			{
				name:     "iframe",
				hasInner: false,
				attributes: []attribute{
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "srcset",
						key:  "srcset",
						typ:  STRING,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "sandbox",
						key:  "sandbox",
						typ:  SPACE_SEPARATED_LIST,
						values: []string{
							"allow-downloads",
							"allow-forms",
							"allow-modals",
							"allow-orientation-lock",
							"allow-pointer-lock",
							"allow-popups",
							"allow-popups-to-escape-sandbox",
							"allow-presentation",
							"allow-same-origin",
							"allow-scripts",
							"allow-top-navigation",
							"allow-top-navigation-by-user-activation",
							"allow-top-navigation-to-custom-protocols",
						},
					},
					{
						name: "allow",
						key:  "allow",
						typ:  STRING,
					},
					{
						name: "allowfullscreen",
						key:  "allowfullscreen",
						typ:  BOOL,
					},
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
					{
						name:   "referrerpolicy",
						key:    "referrerpolicy",
						typ:    STRING,
						values: referrerPolicyValues,
					},
					{
						name:   "loading",
						key:    "loading",
						typ:    STRING,
						values: []string{"lazy", "eager"},
					},
				},
			},
			{
				name:     "embed",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "type",
						key:  "type",
						typ:  STRING,
					},
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "object",
				hasInner: true,
				attributes: []attribute{
					{
						name: "data",
						key:  "data",
						typ:  STRING,
					},
					{
						name: "type",
						key:  "type",
						typ:  STRING,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "form",
						key:  "form",
						typ:  STRING,
					},
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "param",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "value",
						key:  "value",
						typ:  STRING,
					},
				},
			},
			{
				name:     "video",
				hasInner: true,
				attributes: []attribute{
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "crossorigin",
						key:  "crossorigin",
						typ:  STRING,
						values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						name: "poster",
						key:  "poster",
						typ:  STRING,
					},
					{
						name: "preload",
						key:  "preload",
						typ:  STRING,
						values: []string{
							"none",
							"metadata",
							"auto",
						},
					},
					{
						name: "autoplay",
						key:  "autoplay",
						typ:  BOOL,
					},
					{
						name: "playsinline",
						key:  "playsinline",
						typ:  BOOL,
					},
					{
						name: "loop",
						key:  "loop",
						typ:  BOOL,
					},
					{
						name: "muted",
						key:  "muted",
						typ:  BOOL,
					},
					{
						name: "controls",
						key:  "controls",
						typ:  BOOL,
					},
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "audio",
				hasInner: true,
				attributes: []attribute{
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "crossorigin",
						key:  "crossorigin",
						typ:  STRING,
						values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						name: "preload",
						key:  "preload",
						typ:  STRING,
						values: []string{
							"none",
							"metadata",
							"auto",
						},
					},
					{
						name: "autoplay",
						key:  "autoplay",
						typ:  BOOL,
					},
					{
						name: "loop",
						key:  "loop",
						typ:  BOOL,
					},
					{
						name: "muted",
						key:  "muted",
						typ:  BOOL,
					},
					{
						name: "controls",
						key:  "controls",
						typ:  BOOL,
					},
				},
			},
			{
				name:     "track",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "kind",
						key:  "kind",
						typ:  STRING,
						values: []string{
							"subtitles",
							"captions",
							"descriptions",
							"chapters",
							"metadata",
						},
					},
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "srclang",
						key:  "srclang",
						typ:  STRING,
					},
					{
						name: "label",
						key:  "label",
						typ:  STRING,
					},
					{
						name: "default",
						key:  "default",
						typ:  BOOL,
					},
				},
			},
			{
				name:     "map",
				hasInner: true,
				attributes: []attribute{
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
				},
			},
			{
				name:     "area",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "alt",
						key:  "alt",
						typ:  STRING,
					},
					{
						name: "coords",
						key:  "coords",
						typ:  STRING,
					},
					{
						name: "shape",
						key:  "shape",
						typ:  STRING,
						values: []string{
							"circle",
							"default",
							"poly",
							"rect",
						},
					},
					{
						name: "href",
						key:  "href",
						typ:  STRING,
					},
					{
						name: "target",
						key:  "target",
						typ:  STRING,
					},
					{
						name: "download",
						key:  "download",
						typ:  STRING,
					},
					{
						name: "ping",
						key:  "ping",
						typ:  STRING,
					},
					{
						name:   "rel",
						key:    "rel",
						typ:    SPACE_SEPARATED_LIST,
						values: relAttributes(REL_AREA),
					},
					{
						name:   "referrerpolicy",
						key:    "referrerpolicy",
						typ:    STRING,
						values: referrerPolicyValues,
					},
				},
			},
		},
	},
	{
		name: "table",
		elements: []element{
			{
				name:     "table",
				hasInner: false,
			},
			{
				name:     "caption",
				hasInner: true,
			},
			{
				name:     "colgroup",
				hasInner: false,
				attributes: []attribute{
					{
						name: "span",
						key:  "span",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "col",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "span",
						key:  "span",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "tbody",
				hasInner: false,
			},
			{
				name:     "thead",
				hasInner: false,
			},
			{
				name:     "tfoot",
				hasInner: false,
			},
			{
				name:     "tr",
				hasInner: false,
			},
			{
				name:     "td",
				hasInner: true,
				attributes: []attribute{
					{
						name: "colspan",
						key:  "colspan",
						typ:  INTEGER,
					},
					{
						name: "rowspan",
						key:  "rowspan",
						typ:  INTEGER,
					},
					{
						name: "headers",
						key:  "headers",
						typ:  SPACE_SEPARATED_LIST,
					},
				},
			},
			{
				name:     "th",
				hasInner: true,
				attributes: []attribute{
					{
						name: "colspan",
						key:  "colspan",
						typ:  INTEGER,
					},
					{
						name: "rowspan",
						key:  "rowspan",
						typ:  INTEGER,
					},
					{
						name: "headers",
						key:  "headers",
						typ:  SPACE_SEPARATED_LIST,
					},
					{
						name: "scope",
						key:  "scope",
						typ:  STRING,
						values: []string{
							"row",
							"col",
							"rowgroup",
							"colgroup",
							"auto",
						},
					},
					{
						name: "abbr",
						key:  "abbr",
						typ:  STRING,
					},
				},
			},
		},
	},
	{
		name: "form",
		elements: []element{
			{
				name:     "form",
				hasInner: true,
				attributes: []attribute{
					{
						name: "accept-charset",
						key:  "accept-charset",
						typ:  STRING,
					},
					{
						name: "action",
						key:  "action",
						typ:  STRING,
					},
					{
						name:   "autocomplete",
						key:    "autocomplete",
						typ:    STRING,
						values: []string{"on", "off"},
					},
					{
						name: "enctype",
						key:  "enctype",
						typ:  STRING,
						values: []string{
							"application/x-www-form-urlencoded",
							"multipart/form-data",
							"text/plain",
						},
					},
					{
						name: "method",
						key:  "method",
						typ:  STRING,
						values: []string{
							"get",
							"post",
							"dialog",
						},
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "novalidate",
						key:  "novalidate",
						typ:  BOOL,
					},
					{
						name: "target",
						key:  "target",
						typ:  STRING,
					},
					{
						name:   "rel",
						key:    "rel",
						typ:    SPACE_SEPARATED_LIST,
						values: relAttributes(REL_FORM),
					},
				},
			},
			{
				name:     "label",
				hasInner: true,
				attributes: []attribute{
					{
						name: "for",
						key:  "for",
						typ:  STRING,
					},
				},
			},
			{
				name:     "input",
				hasInner: false,
				noEndTag: true,
				attributes: []attribute{
					{
						name: "accept",
						key:  "accept",
						typ:  STRING, /* comma separated tokens */
					},
					{
						name: "alt",
						key:  "alt",
						typ:  STRING,
					},
					{
						name: "autocomplete",
						key:  "autocomplete",
						typ:  STRING, /* TODO */
					},
					{
						name: "checked",
						key:  "checked",
						typ:  BOOL,
					},
					{
						name: "dirname",
						key:  "dirname",
						typ:  STRING,
					},
					{
						name: "disabled",
						key:  "disabled",
						typ:  BOOL,
					},
					{
						name: "form",
						key:  "form",
						typ:  STRING,
					},
					{
						name: "formaction",
						key:  "formaction",
						typ:  STRING,
					},
					{
						name: "formenctype",
						key:  "formenctype",
						typ:  STRING,
						values: []string{
							"application/x-www-form-urlencoded",
							"multipart/form-data",
							"text/plain",
						},
					},
					{
						name: "formmethod",
						key:  "formmethod",
						typ:  STRING,
						values: []string{
							"get",
							"post",
							"dialog",
						},
					},
					{
						name: "formnovalidate",
						key:  "formnovalidate",
						typ:  BOOL,
					},
					{
						name: "formtarget",
						key:  "formtarget",
						typ:  STRING,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
					{
						name: "list",
						key:  "list",
						typ:  STRING,
					},
					{
						name: "max",
						key:  "max",
						typ:  INTEGER,
					},
					{
						name: "maxlength",
						key:  "maxlength",
						typ:  INTEGER,
					},
					{
						name: "min",
						key:  "min",
						typ:  INTEGER,
					},
					{
						name: "minlength",
						key:  "minlength",
						typ:  INTEGER,
					},
					{
						name: "multiple",
						key:  "multiple",
						typ:  BOOL,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "pattern",
						key:  "pattern",
						typ:  STRING,
					},
					{
						name: "placeholder",
						key:  "placeholder",
						typ:  STRING,
					},
					{
						name: "readonly",
						key:  "readonly",
						typ:  BOOL,
					},
					{
						name: "required",
						key:  "required",
						typ:  BOOL,
					},
					{
						name: "size",
						key:  "size",
						typ:  INTEGER,
					},
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "step",
						key:  "step",
						typ:  STRING,
					},
					{
						name: "type",
						key:  "type",
						typ:  STRING,
						values: []string{
							"hidden",
							"text",
							"search",
							"tel",
							"url",
							"email",
							"password",
							"date",
							"month",
							"week",
							"time",
							"datetime-local",
							"number",
							"range",
							"color",
							"checkbox",
							"radio",
							"file",
							"submit",
							"image",
							"reset",
							"button",
						},
					},
					{
						name: "value",
						key:  "value",
						typ:  STRING,
					},
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "button",
				hasInner: true,
				attributes: []attribute{
					{
						name: "disabled",
						key:  "disabled",
						typ:  BOOL,
					},
					{
						name: "form",
						key:  "form",
						typ:  STRING,
					},
					{
						name: "formaction",
						key:  "formaction",
						typ:  STRING,
					},
					{
						name: "formenctype",
						key:  "formenctype",
						typ:  STRING,
						values: []string{
							"application/x-www-form-urlencoded",
							"multipart/form-data",
							"text/plain",
						},
					},
					{
						name: "formmethod",
						key:  "formmethod",
						typ:  STRING,
						values: []string{
							"get",
							"post",
							"dialog",
						},
					},
					{
						name: "formnovalidate",
						key:  "formnovalidate",
						typ:  BOOL,
					},
					{
						name: "formtarget",
						key:  "formtarget",
						typ:  STRING,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "type",
						key:  "type",
						typ:  STRING,
					},
					{
						name: "value",
						key:  "value",
						typ:  STRING,
						values: []string{
							"submit",
							"reset",
							"button",
						},
					},
				},
			},
			{
				name:     "select",
				hasInner: false,
				attributes: []attribute{
					{
						name: "autocomplete",
						key:  "autocomplete",
						typ:  STRING, /* TODO */
					},
					{
						name: "disabled",
						key:  "disabled",
						typ:  BOOL,
					},
					{
						name: "form",
						key:  "form",
						typ:  STRING,
					},
					{
						name: "multiple",
						key:  "multiple",
						typ:  BOOL,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "required",
						key:  "required",
						typ:  BOOL,
					},
					{
						name: "size",
						key:  "size",
						typ:  INTEGER,
					},
				},
			},
			{
				name:     "datalist",
				hasInner: true,
			},
			{
				name:     "optgroup",
				hasInner: false,
				attributes: []attribute{
					{
						name: "disabled",
						key:  "disabled",
						typ:  BOOL,
					},
					{
						name: "label",
						key:  "label",
						typ:  STRING,
					},
				},
			},
			{
				name:     "option",
				hasInner: true,
				attributes: []attribute{
					{
						name: "disabled",
						key:  "disabled",
						typ:  BOOL,
					},
					{
						name: "label",
						key:  "label",
						typ:  STRING,
					},
					{
						name: "selected",
						key:  "selected",
						typ:  BOOL,
					},
					{
						name: "value",
						key:  "value",
						typ:  STRING,
					},
				},
			},
			{
				name:     "textarea",
				hasInner: true,
				attributes: []attribute{
					{
						name: "autocomplete",
						key:  "autocomplete",
						typ:  STRING, /* TODO */
					},
					{
						name: "cols",
						key:  "cols",
						typ:  INTEGER,
					},
					{
						name: "dirname",
						key:  "dirname",
						typ:  STRING,
					},
					{
						name: "disabled",
						key:  "disabled",
						typ:  BOOL,
					},
					{
						name: "form",
						key:  "form",
						typ:  STRING,
					},
					{
						name: "maxlength",
						key:  "maxlength",
						typ:  INTEGER,
					},
					{
						name: "minlength",
						key:  "minlength",
						typ:  INTEGER,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
					{
						name: "placeholder",
						key:  "placeholder",
						typ:  STRING,
					},
					{
						name: "readonly",
						key:  "readonly",
						typ:  BOOL,
					},
					{
						name: "required",
						key:  "required",
						typ:  BOOL,
					},
					{
						name: "rows",
						key:  "rows",
						typ:  INTEGER,
					},
					{
						name: "wrap",
						key:  "wrap",
						typ:  STRING,
						values: []string{
							"soft",
							"hard",
						},
					},
				},
			},
			{
				name:     "output",
				hasInner: true,
				attributes: []attribute{
					{
						name: "for",
						key:  "for",
						typ:  SPACE_SEPARATED_LIST,
					},
					{
						name: "form",
						key:  "form",
						typ:  STRING,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
				},
			},
			{
				name:     "progress",
				hasInner: true,
				attributes: []attribute{
					{
						name: "value",
						key:  "value",
						typ:  FLOAT,
					},
					{
						name: "max",
						key:  "max",
						typ:  FLOAT,
					},
				},
			},
			{
				name:     "meter",
				hasInner: true,
				attributes: []attribute{
					{
						name: "value",
						key:  "value",
						typ:  FLOAT,
					},
					{
						name: "min",
						key:  "min",
						typ:  FLOAT,
					},
					{
						name: "max",
						key:  "max",
						typ:  FLOAT,
					},
					{
						name: "low",
						key:  "low",
						typ:  FLOAT,
					},
					{
						name: "high",
						key:  "high",
						typ:  FLOAT,
					},
					{
						name: "optimum",
						key:  "optimum",
						typ:  FLOAT,
					},
				},
			},
			{
				name:     "fieldset",
				hasInner: true,
				attributes: []attribute{
					{
						name: "disabled",
						key:  "disabled",
						typ:  BOOL,
					},
					{
						name: "form",
						key:  "form",
						typ:  STRING,
					},
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
				},
			},
			{
				name:     "legend",
				hasInner: true,
			},
		},
	},
	{
		name: "interactive",
		elements: []element{
			{
				name:     "details",
				hasInner: false,
				attributes: []attribute{
					{
						name: "open",
						key:  "open",
						typ:  BOOL,
					},
				},
			},
			{
				name:     "summary",
				hasInner: true,
			},
		},
	},
	{
		name: "scripting",
		elements: []element{
			{
				name:     "script",
				hasInner: true,
				attributes: []attribute{
					{
						name: "src",
						key:  "src",
						typ:  STRING,
					},
					{
						name: "type",
						key:  "type",
						typ:  STRING,
					},
					{
						name: "nomodule",
						key:  "nomodule",
						typ:  BOOL,
					},
					{
						name: "async",
						key:  "async",
						typ:  BOOL,
					},
					{
						name: "defer",
						key:  "defer",
						typ:  BOOL,
					},
					{
						name: "crossorigin",
						key:  "crossorigin",
						typ:  STRING,
						values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						name: "integrity",
						key:  "integrity",
						typ:  STRING,
					},
					{
						name:   "referrerpolicy",
						key:    "referrerpolicy",
						typ:    STRING,
						values: referrerPolicyValues,
					},
					{
						name: "blocking",
						key:  "blocking",
						typ:  BOOL,
					},
				},
			},
			{
				name:     "noscript",
				hasInner: true,
			},
			{
				name:     "template",
				hasInner: true,
			},
			{
				name:     "slot",
				hasInner: true,
				attributes: []attribute{
					{
						name: "name",
						key:  "name",
						typ:  STRING,
					},
				},
			},
			{
				name:     "canvas",
				hasInner: true,
				attributes: []attribute{
					{
						name: "width",
						key:  "width",
						typ:  INTEGER,
					},
					{
						name: "height",
						key:  "height",
						typ:  INTEGER,
					},
				},
			},
		},
	},
}

func toGoId(s string) string {
	s = strings.Title(s)
	s = strings.ReplaceAll(s, "-", "")
	return s
}

func toGoValue(s string) string {
	if len(s) == 1 {
		return s
	}
	s = strings.Title(s)
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, "-", "")
	return s
}

func generateCategories(subpath string) {
	for _, category := range categories {
		filename := filepath.Join(subpath, "elements", category.name+".go")
		f := jen.NewFile("elements")
		f.HeaderComment(GENERATED)
		f.ImportName("github.com/feloy/wasmgo/pkg/dom", "dom")

		for _, element := range category.elements {

			generateElement(f, element)
		}
		err := f.Save(filename)
		if err != nil {
			panic(err)
		}
	}
}

func generateElement(f *jen.File, element element) {

	for _, attribute := range element.attributes {
		if len(attribute.values) > 0 {
			generateConstants(f, attribute, element.name)
		}
	}

	if len(element.attributes) > 0 {
		fields := make([]jen.Code, 0, len(element.attributes))
		for _, attribute := range element.attributes {
			field := jen.Id(toGoId(attribute.name)).Id(attribute.typ)
			fields = append(fields, field)
		}
		f.Type().Id(strings.Title(element.name) + "Options").Struct(fields...)
	}

	dict := jen.Dict{
		jen.Id("Name"): jen.Lit(element.name),
	}
	var params []jen.Code
	if element.hasInner {
		params = append(params, jen.Id("inner").String())
		dict[jen.Id("InnerHTML")] = jen.Id("inner")
	}
	if len(element.attributes) > 0 {
		params = append(params, jen.Id("options").Id(strings.Title(element.name)+"Options"))

		stringAttrDict := jen.Dict{}
		for _, attribute := range element.attributes {
			if attribute.typ != STRING {
				continue
			}
			stringAttrDict[jen.Lit(attribute.key)] = jen.Id("options").Dot(toGoValue(attribute.name))
		}
		if len(stringAttrDict) > 0 {
			dict[jen.Id("Attributes")] = jen.Map(jen.String()).String().Values(stringAttrDict)
		}

		sslAttrDict := jen.Dict{}
		for _, attribute := range element.attributes {
			if attribute.typ != SPACE_SEPARATED_LIST {
				continue
			}
			sslAttrDict[jen.Lit(attribute.key)] = jen.Id("options").Dot(strings.Title(attribute.name))
		}
		if len(sslAttrDict) > 0 {
			dict[jen.Id("SSLAttributes")] = jen.Map(jen.String()).Index().String().Values(sslAttrDict)
		}

		boolAttrDict := jen.Dict{}
		for _, attribute := range element.attributes {
			if attribute.typ != BOOL {
				continue
			}
			boolAttrDict[jen.Lit(attribute.key)] = jen.Id("options").Dot(strings.Title(attribute.name))
		}
		if len(boolAttrDict) > 0 {
			dict[jen.Id("BoolAttributes")] = jen.Map(jen.String()).Bool().Values(boolAttrDict)
		}

		intAttrDict := jen.Dict{}
		for _, attribute := range element.attributes {
			if attribute.typ != INTEGER {
				continue
			}
			intAttrDict[jen.Lit(attribute.key)] = jen.Id("options").Dot(strings.Title(attribute.name))
		}
		if len(intAttrDict) > 0 {
			dict[jen.Id("IntAttributes")] = jen.Map(jen.String()).Int().Values(intAttrDict)
		}

		floatAttrDict := jen.Dict{}
		for _, attribute := range element.attributes {
			if attribute.typ != FLOAT {
				continue
			}
			floatAttrDict[jen.Lit(attribute.key)] = jen.Id("options").Dot(strings.Title(attribute.name))
		}
		if len(floatAttrDict) > 0 {
			dict[jen.Id("FloatAttributes")] = jen.Map(jen.String()).Float32().Values(floatAttrDict)
		}
	}

	if element.noEndTag {
		dict[jen.Id("OmitEndTag")] = jen.Id("true")
	}

	f.Func().Id("New"+strings.Title(element.name)).Params(params...).
		Op("*").Qual("github.com/feloy/wasmgo/pkg/dom", "Tag").
		Block(
			jen.Return(jen.Add(jen.Op("&")).Qual("github.com/feloy/wasmgo/pkg/dom", "Tag")).
				Values(dict))
	f.Line()
}

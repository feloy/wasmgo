package specs

var Categories = []Category{
	{
		Name: "metadata",
		Elements: []Element{
			{
				Name:     "head",
				HasInner: false,
			},
			{
				Name:     "title",
				HasInner: true,
			},
			{
				Name:     "base",
				HasInner: false,
				NoEndTag: true,
			},
			{
				Name:     "link",
				HasInner: false,
				NoEndTag: true,
			},
			{
				Name:     "meta",
				HasInner: false,
				NoEndTag: true,
			},
			{
				Name:     "style",
				HasInner: true,
			},
		},
	},
	{
		Name: "section",
		Elements: []Element{
			{
				Name:     "body",
				HasInner: true,
			},
			{
				Name:     "article",
				HasInner: true,
			},
			{
				Name:     "section",
				HasInner: true,
			},
			{
				Name:     "nav",
				HasInner: true,
			},
			{
				Name:     "aside",
				HasInner: true,
			},
			{
				Name:     "h1",
				HasInner: true,
			},
			{
				Name:     "h2",
				HasInner: true,
			},
			{
				Name:     "h3",
				HasInner: true,
			},
			{
				Name:     "h4",
				HasInner: true,
			},
			{
				Name:     "h5",
				HasInner: true,
			},
			{
				Name:     "h6",
				HasInner: true,
			},
			{
				Name:     "hgroup",
				HasInner: true,
			},
			{
				Name:     "header",
				HasInner: true,
			},
			{
				Name:     "footer",
				HasInner: true,
			},
			{
				Name:     "address",
				HasInner: true,
			},
		},
	},
	{
		Name: "grouping",
		Elements: []Element{
			{
				Name:     "p",
				HasInner: true,
			},
			{
				Name:     "hr",
				HasInner: false,
				NoEndTag: true,
			},
			{
				Name:     "pre",
				HasInner: true,
			},
			{
				Name:     "blockquote",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "cite",
						Key:  "cite",
						Typ:  "string",
					},
				},
			},
			{
				Name:     "ol",
				HasInner: false,
				Attributes: []Attribute{
					{
						Name: "reversed",
						Key:  "reversed",
						Typ:  BOOL,
					},
					{
						Name: "start",
						Key:  "start",
						Typ:  INTEGER,
					},
					{
						Name:   "type",
						Key:    "type",
						Typ:    STRING,
						Values: []string{"1", "a", "A", "i", "I"},
					},
				},
			},
			{
				Name:     "ul",
				HasInner: false,
			},
			{
				Name:     "menu",
				HasInner: false,
			},
			{
				Name:     "li",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "value",
						Key:  "value",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "dl",
				HasInner: false,
			},
			{
				Name:     "dt",
				HasInner: true,
			},
			{
				Name:     "dd",
				HasInner: true,
			},
			{
				Name:     "figure",
				HasInner: true,
			},
			{
				Name:     "figcaption",
				HasInner: true,
			},
			{
				Name:     "main",
				HasInner: true,
			},
			{
				Name:     "div",
				HasInner: true,
			},
		},
	},
	{
		Name: "text",
		Elements: []Element{
			{
				Name:     "a",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "href",
						Key:  "href",
						Typ:  STRING,
					},
					{
						Name: "target",
						Key:  "target",
						Typ:  STRING,
					},
					{
						Name: "download",
						Key:  "download",
						Typ:  STRING,
					},
					{
						Name: "ping",
						Key:  "ping",
						Typ:  STRING,
					},
					{
						Name:   "rel",
						Key:    "rel",
						Typ:    SPACE_SEPARATED_LIST,
						Values: relAttributes(REL_A),
					},
					{
						Name: "hreflang",
						Key:  "hreflang",
						Typ:  STRING, /* TODO: BCP 47 */
					},
					{
						Name: "type",
						Key:  "type",
						Typ:  STRING,
					},
					{
						Name:   "referrerpolicy",
						Key:    "referrerpolicy",
						Typ:    STRING,
						Values: referrerPolicyValues,
					},
				},
			},
			{
				Name:     "em",
				HasInner: true,
			},
			{
				Name:     "strong",
				HasInner: true,
			},
			{
				Name:     "small",
				HasInner: true,
			},
			{
				Name:     "s",
				HasInner: true,
			},
			{
				Name:     "cite",
				HasInner: true,
			},
			{
				Name:     "q",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "cite",
						Key:  "cite",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "dfn",
				HasInner: true,
			},
			{
				Name:     "abbr",
				HasInner: true,
			},
			{
				Name:     "ruby",
				HasInner: true,
			},
			{
				Name:     "rt",
				HasInner: true,
			},
			{
				Name:     "rp",
				HasInner: true,
			},
			{
				Name:     "data",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "value",
						Key:  "value",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "time",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "datetime",
						Key:  "datetime",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "code",
				HasInner: true,
			},
			{
				Name:     "var",
				HasInner: true,
			},
			{
				Name:     "samp",
				HasInner: true,
			},
			{
				Name:     "kbd",
				HasInner: true,
			},
			{
				Name:     "sub",
				HasInner: true,
			},
			{
				Name:     "sup",
				HasInner: true,
			},
			{
				Name:     "i",
				HasInner: true,
			},
			{
				Name:     "b",
				HasInner: true,
			},
			{
				Name:     "u",
				HasInner: true,
			},
			{
				Name:     "mark",
				HasInner: true,
			},
			{
				Name:     "bdi",
				HasInner: true,
			},
			{
				Name:     "bdo",
				HasInner: true,
			},
			{
				Name:     "span",
				HasInner: true,
			},
			{
				Name:     "br",
				HasInner: false,
				NoEndTag: true,
			},
			{
				Name:     "wbr",
				HasInner: false,
				NoEndTag: true,
			},
		},
	},
	{
		Name: "edit",
		Elements: []Element{
			{
				Name:     "ins",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "cite",
						Key:  "cite",
						Typ:  STRING,
					},
					{
						Name: "datetime",
						Key:  "datetime",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "del",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "cite",
						Key:  "cite",
						Typ:  STRING,
					},
					{
						Name: "datetime",
						Key:  "datetime",
						Typ:  STRING,
					},
				},
			},
		},
	},
	{
		Name: "embedded",
		Elements: []Element{
			{
				Name:     "picture",
				HasInner: false,
			},
			{
				Name:     "source",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "type",
						Key:  "type",
						Typ:  STRING,
					},
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "srcset",
						Key:  "srcset",
						Typ:  STRING,
					},
					{
						Name: "sizes",
						Key:  "sizes",
						Typ:  STRING,
					},
					{
						Name: "media",
						Key:  "media",
						Typ:  STRING,
					},
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "img",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "alt",
						Key:  "alt",
						Typ:  STRING,
					},
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "srcset",
						Key:  "srcset",
						Typ:  STRING,
					},
					{
						Name: "sizes",
						Key:  "sizes",
						Typ:  STRING,
					},
					{
						Name: "crossorigin",
						Key:  "crossorigin",
						Typ:  STRING,
						Values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						Name: "usemap",
						Key:  "usemap",
						Typ:  STRING,
					},
					{
						Name: "ismap",
						Key:  "ismap",
						Typ:  BOOL,
					},
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
					{
						Name:   "referrerpolicy",
						Key:    "referrerpolicy",
						Typ:    STRING,
						Values: referrerPolicyValues,
					},
					{
						Name:   "decoding",
						Key:    "decoding",
						Typ:    STRING,
						Values: []string{"sync", "async", "auto"},
					},
					{
						Name:   "loading",
						Key:    "loading",
						Typ:    STRING,
						Values: []string{"lazy", "eager"},
					},
				},
			},
			{
				Name:     "iframe",
				HasInner: false,
				Attributes: []Attribute{
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "srcset",
						Key:  "srcset",
						Typ:  STRING,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "sandbox",
						Key:  "sandbox",
						Typ:  SPACE_SEPARATED_LIST,
						Values: []string{
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
						Name: "allow",
						Key:  "allow",
						Typ:  STRING,
					},
					{
						Name: "allowfullscreen",
						Key:  "allowfullscreen",
						Typ:  BOOL,
					},
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
					{
						Name:   "referrerpolicy",
						Key:    "referrerpolicy",
						Typ:    STRING,
						Values: referrerPolicyValues,
					},
					{
						Name:   "loading",
						Key:    "loading",
						Typ:    STRING,
						Values: []string{"lazy", "eager"},
					},
				},
			},
			{
				Name:     "embed",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "type",
						Key:  "type",
						Typ:  STRING,
					},
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "object",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "data",
						Key:  "data",
						Typ:  STRING,
					},
					{
						Name: "type",
						Key:  "type",
						Typ:  STRING,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "form",
						Key:  "form",
						Typ:  STRING,
					},
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "param",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "value",
						Key:  "value",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "video",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "crossorigin",
						Key:  "crossorigin",
						Typ:  STRING,
						Values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						Name: "poster",
						Key:  "poster",
						Typ:  STRING,
					},
					{
						Name: "preload",
						Key:  "preload",
						Typ:  STRING,
						Values: []string{
							"none",
							"metadata",
							"auto",
						},
					},
					{
						Name: "autoplay",
						Key:  "autoplay",
						Typ:  BOOL,
					},
					{
						Name: "playsinline",
						Key:  "playsinline",
						Typ:  BOOL,
					},
					{
						Name: "loop",
						Key:  "loop",
						Typ:  BOOL,
					},
					{
						Name: "muted",
						Key:  "muted",
						Typ:  BOOL,
					},
					{
						Name: "controls",
						Key:  "controls",
						Typ:  BOOL,
					},
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "audio",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "crossorigin",
						Key:  "crossorigin",
						Typ:  STRING,
						Values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						Name: "preload",
						Key:  "preload",
						Typ:  STRING,
						Values: []string{
							"none",
							"metadata",
							"auto",
						},
					},
					{
						Name: "autoplay",
						Key:  "autoplay",
						Typ:  BOOL,
					},
					{
						Name: "loop",
						Key:  "loop",
						Typ:  BOOL,
					},
					{
						Name: "muted",
						Key:  "muted",
						Typ:  BOOL,
					},
					{
						Name: "controls",
						Key:  "controls",
						Typ:  BOOL,
					},
				},
			},
			{
				Name:     "track",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "kind",
						Key:  "kind",
						Typ:  STRING,
						Values: []string{
							"subtitles",
							"captions",
							"descriptions",
							"chapters",
							"metadata",
						},
					},
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "srclang",
						Key:  "srclang",
						Typ:  STRING,
					},
					{
						Name: "label",
						Key:  "label",
						Typ:  STRING,
					},
					{
						Name: "default",
						Key:  "default",
						Typ:  BOOL,
					},
				},
			},
			{
				Name:     "map",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "area",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "alt",
						Key:  "alt",
						Typ:  STRING,
					},
					{
						Name: "coords",
						Key:  "coords",
						Typ:  STRING,
					},
					{
						Name: "shape",
						Key:  "shape",
						Typ:  STRING,
						Values: []string{
							"circle",
							"default",
							"poly",
							"rect",
						},
					},
					{
						Name: "href",
						Key:  "href",
						Typ:  STRING,
					},
					{
						Name: "target",
						Key:  "target",
						Typ:  STRING,
					},
					{
						Name: "download",
						Key:  "download",
						Typ:  STRING,
					},
					{
						Name: "ping",
						Key:  "ping",
						Typ:  STRING,
					},
					{
						Name:   "rel",
						Key:    "rel",
						Typ:    SPACE_SEPARATED_LIST,
						Values: relAttributes(REL_AREA),
					},
					{
						Name:   "referrerpolicy",
						Key:    "referrerpolicy",
						Typ:    STRING,
						Values: referrerPolicyValues,
					},
				},
			},
		},
	},
	{
		Name: "table",
		Elements: []Element{
			{
				Name:     "table",
				HasInner: false,
			},
			{
				Name:     "caption",
				HasInner: true,
			},
			{
				Name:     "colgroup",
				HasInner: false,
				Attributes: []Attribute{
					{
						Name: "span",
						Key:  "span",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "col",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "span",
						Key:  "span",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "tbody",
				HasInner: false,
			},
			{
				Name:     "thead",
				HasInner: false,
			},
			{
				Name:     "tfoot",
				HasInner: false,
			},
			{
				Name:     "tr",
				HasInner: false,
			},
			{
				Name:     "td",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "colspan",
						Key:  "colspan",
						Typ:  INTEGER,
					},
					{
						Name: "rowspan",
						Key:  "rowspan",
						Typ:  INTEGER,
					},
					{
						Name: "headers",
						Key:  "headers",
						Typ:  SPACE_SEPARATED_LIST,
					},
				},
			},
			{
				Name:     "th",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "colspan",
						Key:  "colspan",
						Typ:  INTEGER,
					},
					{
						Name: "rowspan",
						Key:  "rowspan",
						Typ:  INTEGER,
					},
					{
						Name: "headers",
						Key:  "headers",
						Typ:  SPACE_SEPARATED_LIST,
					},
					{
						Name: "scope",
						Key:  "scope",
						Typ:  STRING,
						Values: []string{
							"row",
							"col",
							"rowgroup",
							"colgroup",
							"auto",
						},
					},
					{
						Name: "abbr",
						Key:  "abbr",
						Typ:  STRING,
					},
				},
			},
		},
	},
	{
		Name: "form",
		Elements: []Element{
			{
				Name:     "form",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "accept-charset",
						Key:  "accept-charset",
						Typ:  STRING,
					},
					{
						Name: "action",
						Key:  "action",
						Typ:  STRING,
					},
					{
						Name:   "autocomplete",
						Key:    "autocomplete",
						Typ:    STRING,
						Values: []string{"on", "off"},
					},
					{
						Name: "enctype",
						Key:  "enctype",
						Typ:  STRING,
						Values: []string{
							"application/x-www-form-urlencoded",
							"multipart/form-data",
							"text/plain",
						},
					},
					{
						Name: "method",
						Key:  "method",
						Typ:  STRING,
						Values: []string{
							"get",
							"post",
							"dialog",
						},
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "novalidate",
						Key:  "novalidate",
						Typ:  BOOL,
					},
					{
						Name: "target",
						Key:  "target",
						Typ:  STRING,
					},
					{
						Name:   "rel",
						Key:    "rel",
						Typ:    SPACE_SEPARATED_LIST,
						Values: relAttributes(REL_FORM),
					},
				},
			},
			{
				Name:     "label",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "for",
						Key:  "for",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "input",
				HasInner: false,
				NoEndTag: true,
				Attributes: []Attribute{
					{
						Name: "accept",
						Key:  "accept",
						Typ:  STRING, /* comma separated tokens */
					},
					{
						Name: "alt",
						Key:  "alt",
						Typ:  STRING,
					},
					{
						Name: "autocomplete",
						Key:  "autocomplete",
						Typ:  STRING, /* TODO */
					},
					{
						Name: "checked",
						Key:  "checked",
						Typ:  BOOL,
					},
					{
						Name: "dirname",
						Key:  "dirname",
						Typ:  STRING,
					},
					{
						Name: "disabled",
						Key:  "disabled",
						Typ:  BOOL,
					},
					{
						Name: "form",
						Key:  "form",
						Typ:  STRING,
					},
					{
						Name: "formaction",
						Key:  "formaction",
						Typ:  STRING,
					},
					{
						Name: "formenctype",
						Key:  "formenctype",
						Typ:  STRING,
						Values: []string{
							"application/x-www-form-urlencoded",
							"multipart/form-data",
							"text/plain",
						},
					},
					{
						Name: "formmethod",
						Key:  "formmethod",
						Typ:  STRING,
						Values: []string{
							"get",
							"post",
							"dialog",
						},
					},
					{
						Name: "formnovalidate",
						Key:  "formnovalidate",
						Typ:  BOOL,
					},
					{
						Name: "formtarget",
						Key:  "formtarget",
						Typ:  STRING,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
					{
						Name: "list",
						Key:  "list",
						Typ:  STRING,
					},
					{
						Name: "max",
						Key:  "max",
						Typ:  INTEGER,
					},
					{
						Name: "maxlength",
						Key:  "maxlength",
						Typ:  INTEGER,
					},
					{
						Name: "min",
						Key:  "min",
						Typ:  INTEGER,
					},
					{
						Name: "minlength",
						Key:  "minlength",
						Typ:  INTEGER,
					},
					{
						Name: "multiple",
						Key:  "multiple",
						Typ:  BOOL,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "pattern",
						Key:  "pattern",
						Typ:  STRING,
					},
					{
						Name: "placeholder",
						Key:  "placeholder",
						Typ:  STRING,
					},
					{
						Name: "readonly",
						Key:  "readonly",
						Typ:  BOOL,
					},
					{
						Name: "required",
						Key:  "required",
						Typ:  BOOL,
					},
					{
						Name: "size",
						Key:  "size",
						Typ:  INTEGER,
					},
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "step",
						Key:  "step",
						Typ:  STRING,
					},
					{
						Name: "type",
						Key:  "type",
						Typ:  STRING,
						Values: []string{
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
						Name: "value",
						Key:  "value",
						Typ:  STRING,
					},
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "button",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "disabled",
						Key:  "disabled",
						Typ:  BOOL,
					},
					{
						Name: "form",
						Key:  "form",
						Typ:  STRING,
					},
					{
						Name: "formaction",
						Key:  "formaction",
						Typ:  STRING,
					},
					{
						Name: "formenctype",
						Key:  "formenctype",
						Typ:  STRING,
						Values: []string{
							"application/x-www-form-urlencoded",
							"multipart/form-data",
							"text/plain",
						},
					},
					{
						Name: "formmethod",
						Key:  "formmethod",
						Typ:  STRING,
						Values: []string{
							"get",
							"post",
							"dialog",
						},
					},
					{
						Name: "formnovalidate",
						Key:  "formnovalidate",
						Typ:  BOOL,
					},
					{
						Name: "formtarget",
						Key:  "formtarget",
						Typ:  STRING,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "type",
						Key:  "type",
						Typ:  STRING,
					},
					{
						Name: "value",
						Key:  "value",
						Typ:  STRING,
						Values: []string{
							"submit",
							"reset",
							"button",
						},
					},
				},
			},
			{
				Name:     "select",
				HasInner: false,
				Attributes: []Attribute{
					{
						Name: "autocomplete",
						Key:  "autocomplete",
						Typ:  STRING, /* TODO */
					},
					{
						Name: "disabled",
						Key:  "disabled",
						Typ:  BOOL,
					},
					{
						Name: "form",
						Key:  "form",
						Typ:  STRING,
					},
					{
						Name: "multiple",
						Key:  "multiple",
						Typ:  BOOL,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "required",
						Key:  "required",
						Typ:  BOOL,
					},
					{
						Name: "size",
						Key:  "size",
						Typ:  INTEGER,
					},
				},
			},
			{
				Name:     "datalist",
				HasInner: true,
			},
			{
				Name:     "optgroup",
				HasInner: false,
				Attributes: []Attribute{
					{
						Name: "disabled",
						Key:  "disabled",
						Typ:  BOOL,
					},
					{
						Name: "label",
						Key:  "label",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "option",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "disabled",
						Key:  "disabled",
						Typ:  BOOL,
					},
					{
						Name: "label",
						Key:  "label",
						Typ:  STRING,
					},
					{
						Name: "selected",
						Key:  "selected",
						Typ:  BOOL,
					},
					{
						Name: "value",
						Key:  "value",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "textarea",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "autocomplete",
						Key:  "autocomplete",
						Typ:  STRING, /* TODO */
					},
					{
						Name: "cols",
						Key:  "cols",
						Typ:  INTEGER,
					},
					{
						Name: "dirname",
						Key:  "dirname",
						Typ:  STRING,
					},
					{
						Name: "disabled",
						Key:  "disabled",
						Typ:  BOOL,
					},
					{
						Name: "form",
						Key:  "form",
						Typ:  STRING,
					},
					{
						Name: "maxlength",
						Key:  "maxlength",
						Typ:  INTEGER,
					},
					{
						Name: "minlength",
						Key:  "minlength",
						Typ:  INTEGER,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
					{
						Name: "placeholder",
						Key:  "placeholder",
						Typ:  STRING,
					},
					{
						Name: "readonly",
						Key:  "readonly",
						Typ:  BOOL,
					},
					{
						Name: "required",
						Key:  "required",
						Typ:  BOOL,
					},
					{
						Name: "rows",
						Key:  "rows",
						Typ:  INTEGER,
					},
					{
						Name: "wrap",
						Key:  "wrap",
						Typ:  STRING,
						Values: []string{
							"soft",
							"hard",
						},
					},
				},
			},
			{
				Name:     "output",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "for",
						Key:  "for",
						Typ:  SPACE_SEPARATED_LIST,
					},
					{
						Name: "form",
						Key:  "form",
						Typ:  STRING,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "progress",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "value",
						Key:  "value",
						Typ:  FLOAT,
					},
					{
						Name: "max",
						Key:  "max",
						Typ:  FLOAT,
					},
				},
			},
			{
				Name:     "meter",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "value",
						Key:  "value",
						Typ:  FLOAT,
					},
					{
						Name: "min",
						Key:  "min",
						Typ:  FLOAT,
					},
					{
						Name: "max",
						Key:  "max",
						Typ:  FLOAT,
					},
					{
						Name: "low",
						Key:  "low",
						Typ:  FLOAT,
					},
					{
						Name: "high",
						Key:  "high",
						Typ:  FLOAT,
					},
					{
						Name: "optimum",
						Key:  "optimum",
						Typ:  FLOAT,
					},
				},
			},
			{
				Name:     "fieldset",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "disabled",
						Key:  "disabled",
						Typ:  BOOL,
					},
					{
						Name: "form",
						Key:  "form",
						Typ:  STRING,
					},
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "legend",
				HasInner: true,
			},
		},
	},
	{
		Name: "interactive",
		Elements: []Element{
			{
				Name:     "details",
				HasInner: false,
				Attributes: []Attribute{
					{
						Name: "open",
						Key:  "open",
						Typ:  BOOL,
					},
				},
			},
			{
				Name:     "summary",
				HasInner: true,
			},
		},
	},
	{
		Name: "scripting",
		Elements: []Element{
			{
				Name:     "script",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "src",
						Key:  "src",
						Typ:  STRING,
					},
					{
						Name: "type",
						Key:  "type",
						Typ:  STRING,
					},
					{
						Name: "nomodule",
						Key:  "nomodule",
						Typ:  BOOL,
					},
					{
						Name: "async",
						Key:  "async",
						Typ:  BOOL,
					},
					{
						Name: "defer",
						Key:  "defer",
						Typ:  BOOL,
					},
					{
						Name: "crossorigin",
						Key:  "crossorigin",
						Typ:  STRING,
						Values: []string{
							"anonymous",
							"use-credentials",
						},
					},
					{
						Name: "integrity",
						Key:  "integrity",
						Typ:  STRING,
					},
					{
						Name:   "referrerpolicy",
						Key:    "referrerpolicy",
						Typ:    STRING,
						Values: referrerPolicyValues,
					},
					{
						Name: "blocking",
						Key:  "blocking",
						Typ:  BOOL,
					},
				},
			},
			{
				Name:     "noscript",
				HasInner: true,
			},
			{
				Name:     "template",
				HasInner: true,
			},
			{
				Name:     "slot",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "name",
						Key:  "name",
						Typ:  STRING,
					},
				},
			},
			{
				Name:     "canvas",
				HasInner: true,
				Attributes: []Attribute{
					{
						Name: "width",
						Key:  "width",
						Typ:  INTEGER,
					},
					{
						Name: "height",
						Key:  "height",
						Typ:  INTEGER,
					},
				},
			},
		},
	},
}

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
						name: "referrerpolicy",
						key:  "referrerpolicy",
						typ:  STRING,
						values: []string{
							"no-referrer",
							"no-referrer-when-downgrade",
							"same-origin",
							"origin",
							"strict-origin",
							"origin-when-cross-origin",
							"strict-origin-when-cross-origin",
							"unsafe-url",
						},
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
					},
				},
			},
			{
				name:     "time",
				hasInner: true,
				attributes: []attribute{
					{
						name: "datetime",
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
					},
					{
						name: "datetime",
					},
				},
			},
			{
				name:     "del",
				hasInner: true,
				attributes: []attribute{
					{
						name: "cite",
					},
					{
						name: "datetime",
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
					},
					{
						name: "src",
					},
					{
						name: "srcset",
					},
					{
						name: "sizes",
					},
					{
						name: "media",
					},
					{
						name: "width",
					},
					{
						name: "height",
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
					},
					{
						name: "src",
					},
					{
						name: "srcset",
					},
					{
						name: "sizes",
					},
					{
						name: "crossorigin",
					},
					{
						name: "usemap",
					},
					{
						name: "ismap",
					},
					{
						name: "width",
					},
					{
						name: "height",
					},
					{
						name: "referrerpolicy",
					},
					{
						name: "decoding",
					},
					{
						name: "loading",
					},
				},
			},
			{
				name:     "iframe",
				hasInner: false,
				attributes: []attribute{
					{
						name: "src",
					},
					{
						name: "srcset",
					},
					{
						name: "name",
					},
					{
						name: "sandbox",
					},
					{
						name: "allow",
					},
					{
						name: "allowfullscreen",
					},
					{
						name: "width",
					},
					{
						name: "height",
					},
					{
						name: "referrerpolicy",
					},
					{
						name: "loading",
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
					},
					{
						name: "type",
					},
					{
						name: "width",
					},
					{
						name: "height",
					},
				},
			},
			{
				name:     "object",
				hasInner: true,
				attributes: []attribute{
					{
						name: "data",
					},
					{
						name: "type",
					},
					{
						name: "name",
					},
					{
						name: "form",
					},
					{
						name: "width",
					},
					{
						name: "height",
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
					},
					{
						name: "value",
					},
				},
			},
			{
				name:     "video",
				hasInner: true,
				attributes: []attribute{
					{
						name: "src",
					},
					{
						name: "crossorigin",
					},
					{
						name: "poster",
					},
					{
						name: "preload",
					},
					{
						name: "autoplay",
					},
					{
						name: "playsinline",
					},
					{
						name: "loop",
					},
					{
						name: "muted",
					},
					{
						name: "controls",
					},
					{
						name: "width",
					},
					{
						name: "height",
					},
				},
			},
			{
				name:     "audio",
				hasInner: true,
				attributes: []attribute{
					{
						name: "src",
					},
					{
						name: "crossorigin",
					},
					{
						name: "preload",
					},
					{
						name: "autoplay",
					},
					{
						name: "loop",
					},
					{
						name: "muted",
					},
					{
						name: "controls",
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
					},
					{
						name: "src",
					},
					{
						name: "srclang",
					},
					{
						name: "label",
					},
					{
						name: "default",
					},
				},
			},
			{
				name:     "map",
				hasInner: true,
				attributes: []attribute{
					{
						name: "name",
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
					},
					{
						name: "coords",
					},
					{
						name: "shape",
					},
					{
						name: "href",
					},
					{
						name: "target",
					},
					{
						name: "download",
					},
					{
						name: "ping",
					},
					{
						name: "rel",
					},
					{
						name: "referrerpolicy",
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
					},
					{
						name: "rowspan",
					},
					{
						name: "headers",
					},
				},
			},
			{
				name:     "th",
				hasInner: true,
				attributes: []attribute{
					{
						name: "colspan",
					},
					{
						name: "rowspan",
					},
					{
						name: "headers",
					},
					{
						name: "scope",
					},
					{
						name: "abbr",
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
					},
					{
						name: "action",
					},
					{
						name: "autocomplete",
					},
					{
						name: "enctype",
					},
					{
						name: "method",
					},
					{
						name: "name",
					},
					{
						name: "novalidate",
					},
					{
						name: "target",
					},
					{
						name: "rel",
					},
				},
			},
			{
				name:     "label",
				hasInner: true,
				attributes: []attribute{
					{
						name: "for",
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
					},
					{
						name: "alt",
					},
					{
						name: "autocomplete",
					},
					{
						name: "checked",
					},
					{
						name: "dirname",
					},
					{
						name: "disabled",
					},
					{
						name: "form",
					},
					{
						name: "formaction",
					},
					{
						name: "formenctype",
					},
					{
						name: "formmethod",
					},
					{
						name: "formnovalidate",
					},
					{
						name: "formtarget",
					},
					{
						name: "height",
					},
					{
						name: "list",
					},
					{
						name: "max",
					},
					{
						name: "maxlength",
					},
					{
						name: "min",
					},
					{
						name: "minlength",
					},
					{
						name: "multiple",
					},
					{
						name: "name",
					},
					{
						name: "pattern",
					},
					{
						name: "placeholder",
					},
					{
						name: "readonly",
					},
					{
						name: "required",
					},
					{
						name: "size",
					},
					{
						name: "src",
					},
					{
						name: "step",
					},
					{
						name: "type",
					},
					{
						name: "value",
					},
					{
						name: "width",
					},
				},
			},
			{
				name:     "button",
				hasInner: true,
				attributes: []attribute{
					{
						name: "disabled",
					},
					{
						name: "form",
					},
					{
						name: "formaction",
					},
					{
						name: "formenctype",
					},
					{
						name: "formmethod",
					},
					{
						name: "formnovalidate",
					},
					{
						name: "formtarget",
					},
					{
						name: "name",
					},
					{
						name: "type",
					},
					{
						name: "value",
					},
				},
			},
			{
				name:     "select",
				hasInner: false,
				attributes: []attribute{
					{
						name: "autocomplete",
					},
					{
						name: "disabled",
					},
					{
						name: "form",
					},
					{
						name: "multiple",
					},
					{
						name: "name",
					},
					{
						name: "required",
					},
					{
						name: "size",
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
					},
					{
						name: "label",
					},
				},
			},
			{
				name:     "option",
				hasInner: true,
				attributes: []attribute{
					{
						name: "disabled",
					},
					{
						name: "label",
					},
					{
						name: "selected",
					},
					{
						name: "value",
					},
				},
			},
			{
				name:     "textarea",
				hasInner: true,
				attributes: []attribute{
					{
						name: "autocomplete",
					},
					{
						name: "cols",
					},
					{
						name: "dirname",
					},
					{
						name: "disabled",
					},
					{
						name: "form",
					},
					{
						name: "maxlength",
					},
					{
						name: "minlength",
					},
					{
						name: "name",
					},
					{
						name: "placeholder",
					},
					{
						name: "readonly",
					},
					{
						name: "required",
					},
					{
						name: "rows",
					},
					{
						name: "wrap",
					},
				},
			},
			{
				name:     "output",
				hasInner: true,
				attributes: []attribute{
					{
						name: "for",
					},
					{
						name: "form",
					},
					{
						name: "name",
					},
				},
			},
			{
				name:     "progress",
				hasInner: true,
				attributes: []attribute{
					{
						name: "value",
					},
					{
						name: "max",
					},
				},
			},
			{
				name:     "meter",
				hasInner: true,
				attributes: []attribute{
					{
						name: "value",
					},
					{
						name: "min",
					},
					{
						name: "max",
					},
					{
						name: "low",
					},
					{
						name: "high",
					},
					{
						name: "optimum",
					},
				},
			},
			{
				name:     "fieldset",
				hasInner: true,
				attributes: []attribute{
					{
						name: "disabled",
					},
					{
						name: "form",
					},
					{
						name: "name",
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
					},
					{
						name: "type",
					},
					{
						name: "nomodule",
					},
					{
						name: "async",
					},
					{
						name: "defer",
					},
					{
						name: "crossorigin",
					},
					{
						name: "integrity",
					},
					{
						name: "referrerpolicy",
					},
					{
						name: "blocking",
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
					},
				},
			},
			{
				name:     "canvas",
				hasInner: true,
				attributes: []attribute{
					{
						name: "width",
					},
					{
						name: "height",
					},
				},
			},
		},
	},
}

func toGoId(s string) string {
	return strings.ReplaceAll(strings.Title(s), "-", "")
}

func toGoValue(s string) string {
	if len(s) == 1 {
		return s
	}
	return strings.ReplaceAll(strings.Title(s), "-", "")
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
			stringAttrDict[jen.Lit(attribute.key)] = jen.Id("options").Dot(strings.Title(attribute.name))
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

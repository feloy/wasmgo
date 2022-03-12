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
			},
			{
				name:     "ol",
				hasInner: false,
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
						name: "destination",
						key:  "href",
						typ:  "string",
					},
					{
						name: "relation",
						key:  "rel",
						typ:  "string",
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
			},
			{
				name:     "time",
				hasInner: true,
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
			},
			{
				name:     "del",
				hasInner: true,
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
			},
			{
				name:     "img",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "iframe",
				hasInner: false,
			},
			{
				name:     "embed",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "object",
				hasInner: true,
			},
			{
				name:     "param",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "video",
				hasInner: true,
			},
			{
				name:     "audio",
				hasInner: true,
			},
			{
				name:     "track",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "map",
				hasInner: true,
			},
			{
				name:     "area",
				hasInner: false,
				noEndTag: true,
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
			},
			{
				name:     "col",
				hasInner: false,
				noEndTag: true,
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
			},
			{
				name:     "th",
				hasInner: true,
			},
		},
	},
	{
		name: "form",
		elements: []element{
			{
				name:     "form",
				hasInner: true,
			},
			{
				name:     "label",
				hasInner: true,
			},
			{
				name:     "input",
				hasInner: false,
				noEndTag: true,
			},
			{
				name:     "button",
				hasInner: true,
			},
			{
				name:     "select",
				hasInner: false,
			},
			{
				name:     "datalist",
				hasInner: true,
			},
			{
				name:     "optgroup",
				hasInner: false,
			},
			{
				name:     "option",
				hasInner: true,
			},
			{
				name:     "textarea",
				hasInner: true,
			},
			{
				name:     "output",
				hasInner: true,
			},
			{
				name:     "progress",
				hasInner: true,
			},
			{
				name:     "meter",
				hasInner: true,
			},
			{
				name:     "fieldset",
				hasInner: true,
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
			},
			{
				name:     "canvas",
				hasInner: true,
			},
		},
	},
}

func generateCategories(subpath string) {
	for _, category := range categories {
		filename := filepath.Join(subpath, "elements", category.name+".go")
		f := jen.NewFile("elements")
		f.HeaderComment(GENERATED)
		f.ImportName("github.com/feloy/wasmgo/pkg/dom", "dom")

		for _, element := range category.elements {

			if len(element.attributes) > 0 {
				fields := make([]jen.Code, 0, len(element.attributes))
				for _, attribute := range element.attributes {
					field := jen.Id(strings.Title(attribute.name)).Id(attribute.typ)
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
				attrDict := jen.Dict{}
				for _, attribute := range element.attributes {
					attrDict[jen.Lit(attribute.key)] = jen.Id("options").Dot(strings.Title(attribute.name))
				}
				dict[jen.Id("Attributes")] = jen.Map(jen.String()).String().Values(attrDict)
			}

			if element.noEndTag {
				dict[jen.Id("OmitEndTag")] = jen.Id("true")
			}

			f.Func().Id("New"+strings.Title(element.name)).Params(params...).
				Op("*").Qual("github.com/feloy/wasmgo/pkg/dom", "Tag").
				Block(
					jen.Return(jen.Add(jen.Op("&")).Qual("github.com/feloy/wasmgo/pkg/dom", "Tag")).
						Values(dict))
		}
		err := f.Save(filename)
		if err != nil {
			panic(err)
		}
	}
}

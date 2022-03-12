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
				name: "head",
			},
			{
				name: "title",
			},
			{
				name: "base",
			},
			{
				name: "link",
			},
			{
				name: "meta",
			},
			{
				name: "style",
			},
		},
	},
	{
		name: "section",
		elements: []element{
			{
				name: "body",
			},
			{
				name: "article",
			},
			{
				name: "section",
			},
			{
				name: "nav",
			},
			{
				name: "aside",
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
				name: "hgroup",
			},
			{
				name: "header",
			},
			{
				name: "footer",
			},
			{
				name: "address",
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
				name: "hr",
			},
			{
				name: "pre",
			},
			{
				name: "blockquote",
			},
			{
				name: "ol",
			},
			{
				name: "ul",
			},
			{
				name: "menu",
			},
			{
				name: "li",
			},
			{
				name: "dl",
			},
			{
				name: "dt",
			},
			{
				name: "dd",
			},
			{
				name: "figure",
			},
			{
				name: "figcaption",
			},
			{
				name: "main",
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
				name: "em",
			},
			{
				name: "strong",
			},
			{
				name: "small",
			},
			{
				name: "s",
			},
			{
				name: "cite",
			},
			{
				name: "q",
			},
			{
				name: "dfn",
			},
			{
				name: "abbr",
			},
			{
				name: "ruby",
			},
			{
				name: "rt",
			},
			{
				name: "rp",
			},
			{
				name: "data",
			},
			{
				name: "time",
			},
			{
				name: "code",
			},
			{
				name: "var",
			},
			{
				name: "samp",
			},
			{
				name: "kbd",
			},
			{
				name: "sub",
			},
			{
				name: "sup",
			},
			{
				name: "i",
			},
			{
				name: "b",
			},
			{
				name: "u",
			},
			{
				name: "mark",
			},
			{
				name: "bdi",
			},
			{
				name: "bdo",
			},
			{
				name: "span",
			},
			{
				name: "br",
			},
			{
				name: "wbr",
			},
		},
	},
	{
		name: "edit",
		elements: []element{
			{
				name: "ins",
			},
			{
				name: "del",
			},
		},
	},
	{
		name: "embedded",
		elements: []element{
			{
				name: "picture",
			},
			{
				name: "source",
			},
			{
				name: "img",
			},
			{
				name: "iframe",
			},
			{
				name: "embed",
			},
			{
				name: "object",
			},
			{
				name: "param",
			},
			{
				name: "video",
			},
			{
				name: "audio",
			},
			{
				name: "track",
			},
			{
				name: "map",
			},
			{
				name: "area",
			},
		},
	},
	{
		name: "table",
		elements: []element{
			{
				name: "table",
			},
			{
				name: "caption",
			},
			{
				name: "colgroup",
			},
			{
				name: "col",
			},
			{
				name: "tbody",
			},
			{
				name: "thead",
			},
			{
				name: "tfoot",
			},
			{
				name: "tr",
			},
			{
				name: "td",
			},
			{
				name: "th",
			},
		},
	},
	{
		name: "form",
		elements: []element{
			{
				name: "form",
			},
			{
				name: "label",
			},
			{
				name: "input",
			},
			{
				name:     "button",
				hasInner: true,
			},
			{
				name: "select",
			},
			{
				name: "datalist",
			},
			{
				name: "optgroup",
			},
			{
				name: "option",
			},
			{
				name: "textarea",
			},
			{
				name: "output",
			},
			{
				name: "progress",
			},
			{
				name: "meter",
			},
			{
				name: "fieldset",
			},
			{
				name: "legend",
			},
		},
	},
	{
		name: "interactive",
		elements: []element{
			{
				name: "details",
			},
			{
				name: "summary",
			},
		},
	},
	{
		name: "scripting",
		elements: []element{
			{
				name: "script",
			},
			{
				name: "noscript",
			},
			{
				name: "template",
			},
			{
				name: "slot",
			},
			{
				name: "canvas",
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

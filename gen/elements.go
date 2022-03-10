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
	tag        string
	hasInner   bool
	attributes []attribute
}

type attribute struct {
	name string
	key  string
	typ  string
}

var categories = []category{
	{
		name: "button",
		elements: []element{
			{
				name:     "button",
				tag:      "button",
				hasInner: true,
			},
		},
	},
	{
		name: "headers",
		elements: []element{
			{
				name:     "Header1",
				tag:      "h1",
				hasInner: true,
			},
			{
				name:     "Header2",
				tag:      "h2",
				hasInner: true,
			},
			{
				name:     "Header3",
				tag:      "h3",
				hasInner: true,
			},
			{
				name:     "Header4",
				tag:      "h4",
				hasInner: true,
			},
			{
				name:     "Header5",
				tag:      "h5",
				hasInner: true,
			},
			{
				name:     "Header6",
				tag:      "h6",
				hasInner: true,
			},
		},
	},
	{
		name: "hyperlink",
		elements: []element{
			{
				name:     "hyperlink",
				tag:      "a",
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
		},
	},
	{
		name: "misc",
		elements: []element{
			{
				name:     "div",
				tag:      "div",
				hasInner: false,
			},
			{
				name:     "paragraph",
				tag:      "p",
				hasInner: true,
			},
		},
	},
}

func generateCategories(subpath string) {
	for _, category := range categories {
		filename := filepath.Join(subpath, "elements", category.name+".go")
		f := jen.NewFile("elements")
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
				jen.Id("Name"): jen.Lit(element.tag),
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

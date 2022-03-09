package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

func main() {
	subpath := "../pkg/dom"
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
			fmt.Printf("%v\n", err)
		}
	}
}

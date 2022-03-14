// +build !js

package main

import (
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/feloy/wasmgo/gen/specs"
)

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
	for _, category := range specs.Categories {
		filename := filepath.Join(subpath, "elements", category.Name+".go")
		f := jen.NewFile("elements")
		f.HeaderComment(GENERATED)
		f.ImportName("github.com/feloy/wasmgo/pkg/dom", "dom")

		for _, element := range category.Elements {

			generateElement(f, element)
		}
		err := f.Save(filename)
		if err != nil {
			panic(err)
		}
	}
}

func generateElement(f *jen.File, element specs.Element) {

	for _, attribute := range element.Attributes {
		if len(attribute.Values) > 0 {
			generateConstants(f, attribute, element.Name)
		}
	}

	if len(element.Attributes) > 0 {
		fields := make([]jen.Code, 0, len(element.Attributes))
		for _, attribute := range element.Attributes {
			field := jen.Id(toGoId(attribute.Name)).Id(attribute.Typ)
			fields = append(fields, field)
		}
		f.Type().Id(strings.Title(element.Name) + "Options").Struct(fields...)
	}

	dict := jen.Dict{
		jen.Id("Name"): jen.Lit(element.Name),
	}
	var params []jen.Code
	if element.HasInner {
		params = append(params, jen.Id("inner").String())
		dict[jen.Id("InnerHTML")] = jen.Id("inner")
	}
	if len(element.Attributes) > 0 {
		params = append(params, jen.Id("options").Id(strings.Title(element.Name)+"Options"))

		stringAttrDict := jen.Dict{}
		for _, attribute := range element.Attributes {
			if attribute.Typ != specs.STRING {
				continue
			}
			stringAttrDict[jen.Lit(attribute.Key)] = jen.Id("options").Dot(toGoValue(attribute.Name))
		}
		if len(stringAttrDict) > 0 {
			dict[jen.Id("Attributes")] = jen.Map(jen.String()).String().Values(stringAttrDict)
		}

		sslAttrDict := jen.Dict{}
		for _, attribute := range element.Attributes {
			if attribute.Typ != specs.SPACE_SEPARATED_LIST {
				continue
			}
			sslAttrDict[jen.Lit(attribute.Key)] = jen.Id("options").Dot(strings.Title(attribute.Name))
		}
		if len(sslAttrDict) > 0 {
			dict[jen.Id("SSLAttributes")] = jen.Map(jen.String()).Index().String().Values(sslAttrDict)
		}

		boolAttrDict := jen.Dict{}
		for _, attribute := range element.Attributes {
			if attribute.Typ != specs.BOOL {
				continue
			}
			boolAttrDict[jen.Lit(attribute.Key)] = jen.Id("options").Dot(strings.Title(attribute.Name))
		}
		if len(boolAttrDict) > 0 {
			dict[jen.Id("BoolAttributes")] = jen.Map(jen.String()).Bool().Values(boolAttrDict)
		}

		intAttrDict := jen.Dict{}
		for _, attribute := range element.Attributes {
			if attribute.Typ != specs.INTEGER {
				continue
			}
			intAttrDict[jen.Lit(attribute.Key)] = jen.Id("options").Dot(strings.Title(attribute.Name))
		}
		if len(intAttrDict) > 0 {
			dict[jen.Id("IntAttributes")] = jen.Map(jen.String()).Int().Values(intAttrDict)
		}

		floatAttrDict := jen.Dict{}
		for _, attribute := range element.Attributes {
			if attribute.Typ != specs.FLOAT {
				continue
			}
			floatAttrDict[jen.Lit(attribute.Key)] = jen.Id("options").Dot(strings.Title(attribute.Name))
		}
		if len(floatAttrDict) > 0 {
			dict[jen.Id("FloatAttributes")] = jen.Map(jen.String()).Float32().Values(floatAttrDict)
		}
	}

	if element.NoEndTag {
		dict[jen.Id("OmitEndTag")] = jen.Id("true")
	}

	f.Func().Id("New"+strings.Title(element.Name)).Params(params...).
		Op("*").Qual("github.com/feloy/wasmgo/pkg/dom", "Tag").
		Block(
			jen.Return(jen.Add(jen.Op("&")).Qual("github.com/feloy/wasmgo/pkg/dom", "Tag")).
				Values(dict))
	f.Line()
}

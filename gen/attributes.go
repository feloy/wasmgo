// +build !js

package main

import (
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/feloy/wasmgo/gen/specs"
)

func generateAttributes(subpath string) {
	filename := filepath.Join(subpath, "common_attributes.go")
	f := jen.NewFile("dom")
	f.HeaderComment(GENERATED)

	filename_const := filepath.Join(subpath, "common_attributes_const.go")
	f_const := jen.NewFile("dom")
	f_const.HeaderComment(GENERATED)

	filename_js := filepath.Join(subpath, "common_attributes_js.go")
	f_js := jen.NewFile("dom")
	f_js.HeaderComment("+build js")
	f_js.HeaderComment(GENERATED)

	filename_other := filepath.Join(subpath, "common_attributes_other.go")
	f_other := jen.NewFile("dom")
	f_other.HeaderComment("+build !js")
	f_other.HeaderComment(GENERATED)

	for _, attribute := range specs.CommonAttributes {
		if len(attribute.Values) > 0 {
			generateConstants(f_const, attribute, "")
		}
		if attribute.Key == attribute.Property {
			switch attribute.Typ {
			case specs.BOOL:
				generateBoolAttribute(f, attribute, attribute.Key)
			case specs.INTEGER:
				generateIntAttribute(f, attribute, attribute.Key)
			case specs.SPACE_SEPARATED_LIST:
				generateSpaceSeparatedListAttribute(f, attribute, attribute.Key)
			case specs.STRING:
				generateStringAttribute(f, attribute, attribute.Key)
			}
		} else {
			switch attribute.Typ {
			case specs.BOOL:
				generateBoolAttribute(f_js, attribute, attribute.Property)
				generateBoolAttribute(f_other, attribute, attribute.Key)
			case specs.INTEGER:
				generateIntAttribute(f_js, attribute, attribute.Property)
				generateIntAttribute(f_other, attribute, attribute.Key)
			case specs.SPACE_SEPARATED_LIST:
				generateSpaceSeparatedListAttribute(f_js, attribute, attribute.Property)
				generateSpaceSeparatedListAttribute(f_other, attribute, attribute.Key)
			case specs.STRING:
				generateStringAttribute(f_js, attribute, attribute.Property)
				generateStringAttribute(f_other, attribute, attribute.Key)
			}
		}
	}
	err := f.Save(filename)
	if err != nil {
		panic(err)
	}
	err = f_const.Save(filename_const)
	if err != nil {
		panic(err)
	}
	err = f_js.Save(filename_js)
	if err != nil {
		panic(err)
	}
	err = f_other.Save(filename_other)
	if err != nil {
		panic(err)
	}
}

func generateBoolAttribute(f *jen.File, attribute specs.Attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("With"+strings.Title(attribute.Name)).
		Params(jen.Id(attribute.Name).Bool()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendBoolAttribute").Call(jen.Lit(key), jen.Id(attribute.Name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

func generateIntAttribute(f *jen.File, attribute specs.Attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("With"+strings.Title(attribute.Name)).
		Params(jen.Id(attribute.Name).Int()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendIntAttribute").Call(jen.Lit(key), jen.Id(attribute.Name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

func generateStringAttribute(f *jen.File, attribute specs.Attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("With"+strings.Title(attribute.Name)).
		Params(jen.Id(attribute.Name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendAttribute").Call(jen.Lit(key), jen.Id(attribute.Name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

func generateConstants(f *jen.File, attribute specs.Attribute, prefix string) {
	consts := make([]jen.Code, 0, len(attribute.Values))
	for _, value := range attribute.Values {
		var id string
		if len(prefix) > 0 {
			id = toGoId(prefix) + "_"
		}
		id += toGoId(attribute.Name) + "_" + toGoValue(value)
		cnst := jen.Id(id).Op("=").Lit(value)
		consts = append(consts, cnst)
	}
	f.Const().Defs(consts...)
	f.Line()
}

func generateSpaceSeparatedListAttribute(f *jen.File, attribute specs.Attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("Append"+strings.Title(attribute.Name)).
		Params(jen.Id(attribute.Name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendSSLAttribute").Call(jen.Lit(key), jen.Id(attribute.Name)),
			jen.Return(jen.Id("o")),
		)

	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("Prepend"+strings.Title(attribute.Name)).
		Params(jen.Id(attribute.Name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("prependSSLAttribute").Call(jen.Lit(key), jen.Id(attribute.Name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

package main

import (
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

const (
	SPACE_SEPARATED_LIST = "space_separated_list"
	STRING               = "string"
)

var commonAttributes = []attribute{
	{
		name: "class",
		key:  "class",
		typ:  SPACE_SEPARATED_LIST,
	},
	{
		name: "id",
		key:  "id",
		typ:  STRING,
	},
}

/*
accesskey 			SSL
autocapitalize		ENUM
autofocus			BOOL
class				SSL
contenteditable		ENUM(true/false)
dir					ENUM(auto/ltr/rtl)
draggable			ENUM(true/false)
enterkeyhint		ENUM
hidden				BOOL
id					String
inputmode			ENUM
is					String
itemid				String
itemprop			SSL
itemref				SSL
itemscope			BOOL
itemtype			SSL
lang				ENUM(bcp 47)
nonce				String
role				ENUM(ARIA role)
slot				String
spellcheck			ENUM(true/false)
style				String
tabindex			Int
title				String
translate			ENUM(yes/no)
*/

func generateAttributes(subpath string) {
	filename := filepath.Join(subpath, "common_attributes.go")
	f := jen.NewFile("dom")

	for _, attribute := range commonAttributes {
		switch attribute.typ {
		case STRING:
			generateStringAttribute(f, attribute)
		case SPACE_SEPARATED_LIST:
			generateSpaceSeparatedListAttribute(f, attribute)
		}
	}
	err := f.Save(filename)
	if err != nil {
		panic(err)
	}

}

func generateStringAttribute(f *jen.File, attribute attribute) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("With"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendAttribute").Call(jen.Lit(attribute.key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

func generateSpaceSeparatedListAttribute(f *jen.File, attribute attribute) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("Append"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendSSLAttribute").Call(jen.Lit(attribute.key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("Prepend"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("prependSSLAttribute").Call(jen.Lit(attribute.key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

package main

import (
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

const (
	BOOL                 = "bool"
	ENUM                 = "enum"
	INTEGER              = "integer"
	SPACE_SEPARATED_LIST = "space_separated_list"
	STRING               = "string"
)

var commonAttributes = []attribute{
	{
		name:     "accesskey",
		property: "accessKey",
		key:      "accesskey",
		typ:      SPACE_SEPARATED_LIST,
	},
	{
		name:     "autocapitalize",
		property: "autocapitalize",
		key:      "autocapitalize",
		typ:      ENUM,
		values: []string{
			"characters",
			"none",
			"off",
			"on",
			"sentences",
			"words",
		},
	},
	{
		name:     "autofocus",
		property: "autofocus",
		key:      "autofocus",
		typ:      BOOL,
	},
	{
		name:     "class",
		property: "className",
		key:      "class",
		typ:      SPACE_SEPARATED_LIST,
	},
	{
		name:     "contenteditable",
		property: "contentEditable",
		key:      "contenteditable",
		typ:      ENUM,
		values: []string{
			"true",
			"false",
		},
	},
	{
		name:     "dir",
		property: "dir",
		key:      "dir",
		typ:      ENUM,
		values: []string{
			"auto",
			"ltr",
			"rtl",
		},
	},
	{
		name:     "draggable",
		property: "draggable",
		key:      "draggable",
		typ:      ENUM,
		values: []string{
			"true",
			"false",
		},
	},
	{
		name:     "enterkeyhint",
		property: "enterKeyHint",
		key:      "enterkeyhint",
		typ:      ENUM,
		values: []string{
			"done",
			"enter",
			"go",
			"next",
			"previous",
			"search",
			"send",
		},
	},
	{
		name:     "hidden",
		property: "hidden",
		key:      "hidden",
		typ:      BOOL,
	},
	{
		name:     "id",
		property: "id",
		key:      "id",
		typ:      STRING,
	},
	{
		name:     "inputmode",
		property: "inputMode",
		key:      "inputmode",
		typ:      ENUM,
		values: []string{
			"decimal",
			"email",
			"none",
			"numeric",
			"search",
			"tel",
			"text",
			"url",
		},
	},
	{
		name:     "is",
		property: "is",
		key:      "is",
		typ:      STRING,
	},
	{
		name:     "itemid",
		property: "itemid",
		key:      "itemid",
		typ:      STRING,
	},
	{
		name:     "itemprop",
		property: "itemprop",
		key:      "itemprop",
		typ:      SPACE_SEPARATED_LIST,
	},
	{
		name:     "itemref",
		property: "itemref",
		key:      "itemref",
		typ:      SPACE_SEPARATED_LIST,
	},
	{
		name:     "itemscope",
		property: "itemscope",
		key:      "itemscope",
		typ:      BOOL,
	},
	{
		name:     "itemtype",
		property: "itemtype",
		key:      "itemtype",
		typ:      SPACE_SEPARATED_LIST,
	},
	{
		name:     "lang",
		property: "lang",
		key:      "lang",
		typ:      STRING,
	},
	{
		name:     "nonce",
		property: "nonce",
		key:      "nonce",
		typ:      STRING,
	},
	{
		name:     "role",
		property: "role",
		key:      "role",
		typ:      ENUM,
		values: []string{
			"alert",
			"alertdialog",
			"application",
			"article",
			"associationlist",
			"associationlistitemkey",
			"associationlistitemvalue",
			"banner",
			"blockquote",
			"caption",
			"cell",
			"columnheader",
			"command",
			"comment",
			"complementary",
			"composite",
			"contentinfo",
			"definition",
			"deletion",
			"dialog",
			"directory",
			"document",
			"emphasis",
			"feed",
			"figure",
			"form",
			"generic",
			"gridcell",
			"group",
			"heading",
			"img",
			"input",
			"insertion",
			"label",
			"legend",
			"list",
			"listitem",
			"log",
			"main",
			"mark",
			"marquee",
			"math",
			"meter",
			"navigation",
			"none",
			"note",
			"paragraph",
			"presentation",
			"progressbar",
			"region",
			"row",
			"rowgroup",
			"rowheader",
			"scrollbar",
			"search",
			"separator",
			"status",
			"strong",
			"subscript",
			"suggestion",
			"superscript",
			"tab",
			"table",
			"term",
			"time",
			"timer",
			"toolbar",
			"tooltip",
		},
	},
	{
		name:     "slot",
		property: "slot",
		key:      "slot",
		typ:      STRING,
	},
	{
		name:     "spellcheck",
		property: "spellcheck",
		key:      "spellcheck",
		typ:      ENUM,
		values: []string{
			"true",
			"false",
		},
	},
	{
		name:     "style",
		property: "style",
		key:      "style",
		typ:      STRING,
	},
	{
		name:     "title",
		property: "title",
		key:      "title",
		typ:      STRING,
	},
	{
		name:     "tabindex",
		property: "tabIndex",
		key:      "tabindex",
		typ:      INTEGER,
	},
	//	{
	//		name: "translate",
	//		property: "translate",
	//		key:  "translate",
	//		typ:  ENUM,
	//		values: []string{
	//			"yes",
	//			"no",
	//		},
	//	},
}

/*
// accesskey 			SSL
// autocapitalize		ENUM
// autofocus			BOOL
// class				SSL
// contenteditable		ENUM(true/false)
// dir					ENUM(auto/ltr/rtl)
// draggable			ENUM(true/false)
// enterkeyhint		ENUM(done,enter,go,next,previous,search,send)
// hidden				BOOL
// id					String
// inputmode			ENUM
// is					String
// itemid				String
// itemprop			SSL
// itemref				SSL
// itemscope			BOOL
// itemtype			SSL
// lang				ENUM(bcp 47)
// nonce				String
// role				ENUM(ARIA role)
// slot				String
// spellcheck			ENUM(true/false)
// style				String
// tabIndex			Int
// title				String
translate			ENUM(yes/no)
*/

func generateAttributes(subpath string) {
	filename := filepath.Join(subpath, "common_attributes.go")
	filename_const := filepath.Join(subpath, "common_attributes_const.go")
	filename_js := filepath.Join(subpath, "common_attributes_js.go")
	filename_other := filepath.Join(subpath, "common_attributes_other.go")
	f := jen.NewFile("dom")
	f_const := jen.NewFile("dom")
	f_js := jen.NewFile("dom")
	f_js.HeaderComment("+build js")
	f_other := jen.NewFile("dom")
	f_other.HeaderComment("+build !js")

	for _, attribute := range commonAttributes {
		if attribute.key == attribute.property {
			switch attribute.typ {
			case BOOL:
				generateBoolAttribute(f, attribute, attribute.key)
			case ENUM:
				generateConstants(f_const, attribute)
				generateStringAttribute(f, attribute, attribute.key)
			case INTEGER:
				generateIntAttribute(f, attribute, attribute.key)
			case SPACE_SEPARATED_LIST:
				generateSpaceSeparatedListAttribute(f, attribute, attribute.key)
			case STRING:
				generateStringAttribute(f, attribute, attribute.key)
			}
		} else {
			switch attribute.typ {
			case BOOL:
				generateBoolAttribute(f_js, attribute, attribute.property)
				generateBoolAttribute(f_other, attribute, attribute.key)
			case ENUM:
				generateConstants(f_const, attribute)
				generateStringAttribute(f_js, attribute, attribute.property)
				generateStringAttribute(f_other, attribute, attribute.key)
			case INTEGER:
				generateIntAttribute(f_js, attribute, attribute.property)
				generateIntAttribute(f_other, attribute, attribute.key)
			case SPACE_SEPARATED_LIST:
				generateSpaceSeparatedListAttribute(f_js, attribute, attribute.property)
				generateSpaceSeparatedListAttribute(f_other, attribute, attribute.key)
			case STRING:
				generateStringAttribute(f_js, attribute, attribute.property)
				generateStringAttribute(f_other, attribute, attribute.key)
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

func generateBoolAttribute(f *jen.File, attribute attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("With"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).Bool()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendBoolAttribute").Call(jen.Lit(key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

func generateIntAttribute(f *jen.File, attribute attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("With"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).Int()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendIntAttribute").Call(jen.Lit(key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

func generateStringAttribute(f *jen.File, attribute attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("With"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendAttribute").Call(jen.Lit(key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

func generateConstants(f *jen.File, attribute attribute) {
	consts := make([]jen.Code, 0, len(attribute.values))
	for _, value := range attribute.values {
		cnst := jen.Id(strings.Title(attribute.name) + strings.Title(value)).Op("=").Lit(value)
		consts = append(consts, cnst)
	}
	f.Const().Defs(consts...)
	f.Line()
}

func generateSpaceSeparatedListAttribute(f *jen.File, attribute attribute, key string) {
	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("Append"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("appendSSLAttribute").Call(jen.Lit(key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Func().
		Params(jen.Id("o").Op("*").Id("Tag")).
		Id("Prepend"+strings.Title(attribute.name)).
		Params(jen.Id(attribute.name).String()).
		Op("*").Id("Tag").
		Block(
			jen.Id("o").Dot("prependSSLAttribute").Call(jen.Lit(key), jen.Id(attribute.name)),
			jen.Return(jen.Id("o")),
		)

	f.Line()
}

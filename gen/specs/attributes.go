package specs

var CommonAttributes = []Attribute{
	{
		Name:     "accesskey",
		Property: "accessKey",
		Key:      "accesskey",
		Typ:      SPACE_SEPARATED_LIST,
	},
	{
		Name:     "autocapitalize",
		Property: "autocapitalize",
		Key:      "autocapitalize",
		Typ:      STRING,
		Values: []string{
			"characters",
			"none",
			"off",
			"on",
			"sentences",
			"words",
		},
	},
	{
		Name:     "autofocus",
		Property: "autofocus",
		Key:      "autofocus",
		Typ:      BOOL,
	},
	{
		Name:     "class",
		Property: "className",
		Key:      "class",
		Typ:      SPACE_SEPARATED_LIST,
	},
	{
		Name:     "contenteditable",
		Property: "contentEditable",
		Key:      "contenteditable",
		Typ:      STRING,
		Values: []string{
			"true",
			"false",
		},
	},
	{
		Name:     "dir",
		Property: "dir",
		Key:      "dir",
		Typ:      STRING,
		Values: []string{
			"auto",
			"ltr",
			"rtl",
		},
	},
	{
		Name:     "draggable",
		Property: "draggable",
		Key:      "draggable",
		Typ:      STRING,
		Values: []string{
			"true",
			"false",
		},
	},
	{
		Name:     "enterkeyhint",
		Property: "enterKeyHint",
		Key:      "enterkeyhint",
		Typ:      STRING,
		Values: []string{
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
		Name:     "hidden",
		Property: "hidden",
		Key:      "hidden",
		Typ:      BOOL,
	},
	{
		Name:     "id",
		Property: "id",
		Key:      "id",
		Typ:      STRING,
	},
	{
		Name:     "inputmode",
		Property: "inputMode",
		Key:      "inputmode",
		Typ:      STRING,
		Values: []string{
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
		Name:     "is",
		Property: "is",
		Key:      "is",
		Typ:      STRING,
	},
	{
		Name:     "itemid",
		Property: "itemid",
		Key:      "itemid",
		Typ:      STRING,
	},
	{
		Name:     "itemprop",
		Property: "itemprop",
		Key:      "itemprop",
		Typ:      SPACE_SEPARATED_LIST,
	},
	{
		Name:     "itemref",
		Property: "itemref",
		Key:      "itemref",
		Typ:      SPACE_SEPARATED_LIST,
	},
	{
		Name:     "itemscope",
		Property: "itemscope",
		Key:      "itemscope",
		Typ:      BOOL,
	},
	{
		Name:     "itemtype",
		Property: "itemtype",
		Key:      "itemtype",
		Typ:      SPACE_SEPARATED_LIST,
	},
	{
		Name:     "lang",
		Property: "lang",
		Key:      "lang",
		Typ:      STRING,
	},
	{
		Name:     "nonce",
		Property: "nonce",
		Key:      "nonce",
		Typ:      STRING,
	},
	{
		Name:     "role",
		Property: "role",
		Key:      "role",
		Typ:      STRING,
		Values: []string{
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
		Name:     "slot",
		Property: "slot",
		Key:      "slot",
		Typ:      STRING,
	},
	{
		Name:     "spellcheck",
		Property: "spellcheck",
		Key:      "spellcheck",
		Typ:      STRING,
		Values: []string{
			"true",
			"false",
		},
	},
	{
		Name:     "style",
		Property: "style",
		Key:      "style",
		Typ:      STRING,
	},
	{
		Name:     "title",
		Property: "title",
		Key:      "title",
		Typ:      STRING,
	},
	{
		Name:     "tabindex",
		Property: "tabIndex",
		Key:      "tabindex",
		Typ:      INTEGER,
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

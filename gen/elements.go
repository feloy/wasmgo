package main

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

// +build !js

package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/feloy/wasmgo/gen/specs"
)

func TestDoubleElements(t *testing.T) {
	elements := map[string]string{}
	for _, category := range specs.Categories {
		for _, element := range category.Elements {
			if firstCategory, found := elements[element.Name]; found {
				t.Errorf("element %q in category %q already found in category %q", element.Name, category.Name, firstCategory)
			}
			elements[element.Name] = category.Name
		}
	}
}

func Test_generateElement(t *testing.T) {
	type args struct {
		element specs.Element
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "an element without innerHTML",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
				},
			},
			want: `func NewAname() *dom.Tag {
	return &dom.Tag{Name: "aname"}
}
`,
		},
		{
			name: "an element with innerHTML",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: true,
				},
			},
			want: `func NewAname(inner string) *dom.Tag {
	return &dom.Tag{
		InnerHTML: inner,
		Name:      "aname",
	}
}
`,
		},
		{
			name: "an element with no end tag",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
					NoEndTag: true,
				},
			},
			want: `func NewAname() *dom.Tag {
	return &dom.Tag{
		Name:       "aname",
		OmitEndTag: true,
	}
}
`,
		},
		{
			name: "an element with string attributes",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
					Attributes: []specs.Attribute{
						{
							Name: "anattrname",
							Key:  "anattrkey",
							Typ:  specs.STRING,
						},
						{
							Name: "anotherattrname",
							Key:  "anotherattrkey",
							Typ:  specs.STRING,
						},
					},
				},
			},
			want: `type AnameOptions struct {
	Anattrname      string
	Anotherattrname string
}

func NewAname(options AnameOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"anattrkey":      options.Anattrname,
			"anotherattrkey": options.Anotherattrname,
		},
		Name: "aname",
	}
}
`,
		},
		{
			name: "an element with SSL attributes",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
					Attributes: []specs.Attribute{
						{
							Name: "anattrname",
							Key:  "anattrkey",
							Typ:  specs.SPACE_SEPARATED_LIST,
						},
						{
							Name: "anotherattrname",
							Key:  "anotherattrkey",
							Typ:  specs.SPACE_SEPARATED_LIST,
						},
					},
				},
			},
			want: `type AnameOptions struct {
	Anattrname      []string
	Anotherattrname []string
}

func NewAname(options AnameOptions) *dom.Tag {
	return &dom.Tag{
		Name: "aname",
		SSLAttributes: map[string][]string{
			"anattrkey":      options.Anattrname,
			"anotherattrkey": options.Anotherattrname,
		},
	}
}
`,
		},
		{
			name: "an element with bool attributes",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
					Attributes: []specs.Attribute{
						{
							Name: "anattrname",
							Key:  "anattrkey",
							Typ:  specs.BOOL,
						},
						{
							Name: "anotherattrname",
							Key:  "anotherattrkey",
							Typ:  specs.BOOL,
						},
					},
				},
			},
			want: `type AnameOptions struct {
	Anattrname      bool
	Anotherattrname bool
}

func NewAname(options AnameOptions) *dom.Tag {
	return &dom.Tag{
		BoolAttributes: map[string]bool{
			"anattrkey":      options.Anattrname,
			"anotherattrkey": options.Anotherattrname,
		},
		Name: "aname",
	}
}
`,
		},
		{
			name: "an element with int attributes",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
					Attributes: []specs.Attribute{
						{
							Name: "anattrname",
							Key:  "anattrkey",
							Typ:  specs.INTEGER,
						},
						{
							Name: "anotherattrname",
							Key:  "anotherattrkey",
							Typ:  specs.INTEGER,
						},
					},
				},
			},
			want: `type AnameOptions struct {
	Anattrname      int
	Anotherattrname int
}

func NewAname(options AnameOptions) *dom.Tag {
	return &dom.Tag{
		IntAttributes: map[string]int{
			"anattrkey":      options.Anattrname,
			"anotherattrkey": options.Anotherattrname,
		},
		Name: "aname",
	}
}
`,
		},
		{
			name: "an element with float attributes",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
					Attributes: []specs.Attribute{
						{
							Name: "anattrname",
							Key:  "anattrkey",
							Typ:  specs.FLOAT,
						},
						{
							Name: "anotherattrname",
							Key:  "anotherattrkey",
							Typ:  specs.FLOAT,
						},
					},
				},
			},
			want: `type AnameOptions struct {
	Anattrname      float32
	Anotherattrname float32
}

func NewAname(options AnameOptions) *dom.Tag {
	return &dom.Tag{
		FloatAttributes: map[string]float32{
			"anattrkey":      options.Anattrname,
			"anotherattrkey": options.Anotherattrname,
		},
		Name: "aname",
	}
}
`,
		},
		{
			name: "an element with string attribute with values",
			args: args{
				element: specs.Element{
					Name:     "aname",
					HasInner: false,
					Attributes: []specs.Attribute{
						{
							Name:   "anattrname",
							Key:    "anattrkey",
							Typ:    specs.STRING,
							Values: []string{"A", "a"},
						},
						{
							Name: "anotherattrname",
							Key:  "anotherattrkey",
							Typ:  specs.STRING,
						},
					},
				},
			},
			want: `const (
	Aname_Anattrname_A = "A"
	Aname_Anattrname_a = "a"
)

type AnameOptions struct {
	Anattrname      string
	Anotherattrname string
}

func NewAname(options AnameOptions) *dom.Tag {
	return &dom.Tag{
		Attributes: map[string]string{
			"anattrkey":      options.Anattrname,
			"anotherattrkey": options.Anotherattrname,
		},
		Name: "aname",
	}
}
`,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prefix := `package mypackage

import dom "github.com/feloy/wasmgo/pkg/dom"

`
			f := jen.NewFile("mypackage")
			generateElement(f, tt.args.element)
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			wanted := prefix + tt.want
			if result != wanted {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", wanted, result)
			}
		})
	}
}

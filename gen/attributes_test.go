// +build !js

package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/feloy/wasmgo/gen/specs"
)

func Test_generateBoolAttribute(t *testing.T) {
	type args struct {
		attribute specs.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a bool attribute",
			args: args{
				attribute: specs.Attribute{
					Name: "aname",
					Key:  "akey",
					Typ:  specs.BOOL,
				},
			},
			want: `package mypackage

func (o *Tag) WithAname(aname bool) *Tag {
	o.appendBoolAttribute("akey", aname)
	return o
}
`,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := jen.NewFile("mypackage")
			generateBoolAttribute(f, tt.args.attribute, tt.args.attribute.Key)
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			if result != tt.want {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", tt.want, result)
			}
		})
	}
}

func Test_generateIntAttribute(t *testing.T) {
	type args struct {
		attribute specs.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "an int attribute",
			args: args{
				attribute: specs.Attribute{
					Name: "aname",
					Key:  "akey",
					Typ:  specs.INTEGER,
				},
			},
			want: `package mypackage

func (o *Tag) WithAname(aname int) *Tag {
	o.appendIntAttribute("akey", aname)
	return o
}
`,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := jen.NewFile("mypackage")
			generateIntAttribute(f, tt.args.attribute, tt.args.attribute.Key)
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			if result != tt.want {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", tt.want, result)
			}
		})
	}
}

func Test_generateStringAttribute(t *testing.T) {
	type args struct {
		attribute specs.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a string attribute",
			args: args{
				attribute: specs.Attribute{
					Name: "aname",
					Key:  "akey",
					Typ:  specs.STRING,
				},
			},
			want: `package mypackage

func (o *Tag) WithAname(aname string) *Tag {
	o.appendAttribute("akey", aname)
	return o
}
`,
		},
	}
	// TODO: Add test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := jen.NewFile("mypackage")
			generateStringAttribute(f, tt.args.attribute, tt.args.attribute.Key)
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			if result != tt.want {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", tt.want, result)
			}
		})
	}
}

func Test_generateSpaceSeparatedListAttribute(t *testing.T) {
	type args struct {
		attribute specs.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a space separated list attribute",
			args: args{
				attribute: specs.Attribute{
					Name: "aname",
					Key:  "akey",
					Typ:  specs.SPACE_SEPARATED_LIST,
				},
			},
			want: `package mypackage

func (o *Tag) AppendAname(aname string) *Tag {
	o.appendSSLAttribute("akey", aname)
	return o
}
func (o *Tag) PrependAname(aname string) *Tag {
	o.prependSSLAttribute("akey", aname)
	return o
}
`,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := jen.NewFile("mypackage")
			generateSpaceSeparatedListAttribute(f, tt.args.attribute, tt.args.attribute.Key)
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			if result != tt.want {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", tt.want, result)
			}
		})
	}
}

func Test_generateConstants(t *testing.T) {
	type args struct {
		attribute specs.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "constants",
			args: args{
				attribute: specs.Attribute{
					Name: "anattr",
					Values: []string{
						"value1",
						"value2",
						"value3",
					},
				},
			},
			want: `package mypackage

const (
	Anattr_Value1 = "value1"
	Anattr_Value2 = "value2"
	Anattr_Value3 = "value3"
)
`,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := jen.NewFile("mypackage")
			generateConstants(f, tt.args.attribute, "")
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			if result != tt.want {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", tt.want, result)
			}
		})
	}
}

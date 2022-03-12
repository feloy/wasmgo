// +build !js

package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
)

func Test_generateBoolAttribute(t *testing.T) {
	type args struct {
		attribute attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a bool attribute",
			args: args{
				attribute: attribute{
					name: "aname",
					key:  "akey",
					typ:  BOOL,
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
			generateBoolAttribute(f, tt.args.attribute, tt.args.attribute.key)
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
		attribute attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "an int attribute",
			args: args{
				attribute: attribute{
					name: "aname",
					key:  "akey",
					typ:  INTEGER,
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
			generateIntAttribute(f, tt.args.attribute, tt.args.attribute.key)
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
		attribute attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a string attribute",
			args: args{
				attribute: attribute{
					name: "aname",
					key:  "akey",
					typ:  STRING,
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
			generateStringAttribute(f, tt.args.attribute, tt.args.attribute.key)
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
		attribute attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a space separated list attribute",
			args: args{
				attribute: attribute{
					name: "aname",
					key:  "akey",
					typ:  SPACE_SEPARATED_LIST,
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
			generateSpaceSeparatedListAttribute(f, tt.args.attribute, tt.args.attribute.key)
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
		attribute attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "constants",
			args: args{
				attribute: attribute{
					name: "anattr",
					values: []string{
						"value1",
						"value2",
						"value3",
					},
				},
			},
			want: `package mypackage

const (
	AnattrValue1 = "value1"
	AnattrValue2 = "value2"
	AnattrValue3 = "value3"
)
`,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := jen.NewFile("mypackage")
			generateConstants(f, tt.args.attribute)
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			if result != tt.want {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", tt.want, result)
			}
		})
	}
}

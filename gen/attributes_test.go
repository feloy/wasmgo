package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
)

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
			generateStringAttribute(f, tt.args.attribute)
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
			generateSpaceSeparatedListAttribute(f, tt.args.attribute)
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "%#v", f)
			result := b.String()
			if result != tt.want {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", tt.want, result)
			}
		})
	}
}

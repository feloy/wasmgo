// +build !js

package main

import "testing"

func TestDoubleElements(t *testing.T) {
	elements := map[string]string{}
	for _, category := range categories {
		for _, element := range category.elements {
			if firstCategory, found := elements[element.name]; found {
				t.Errorf("element %q in category %q already found in category %q", element.name, category.name, firstCategory)
			}
			elements[element.name] = category.name
		}
	}
}

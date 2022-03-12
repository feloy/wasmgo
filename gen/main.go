// +build !js

package main

func main() {
	subpath := "../pkg/dom"

	generateCategories(subpath)
	generateAttributes(subpath)
}

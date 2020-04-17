// This is a fake package manager that does nothing. It is only an excercise.
// It only prints what it would do using a static list.
package main

import "log"

// Pkg defines the package name, version, and dependencies
type Pkg struct {
	name         string
	version      string
	dependencies []Pkg
}

// MakeFakeList returns a list of fake packages and deps
func MakeFakeList() *Pkg {
	pkgs := Pkg{
		name:    "golang",
		version: "2:1.12~1ubuntu1",
		dependencies: []Pkg{
			{name: "golang-doc", version: "2:1.12~1ubuntu1"},
			{name: "golang-src", version: "2:1.12~1ubuntu1"},
			{
				name:         "golang-go",
				version:      "2:1.12~1ubuntu1",
				dependencies: []Pkg{{name: "gcc", version: "4:9.2.1-3.1ubuntu1"}}},
		},
	}

	return &pkgs
}

func main() {
	pkgs := MakeFakeList()

	log.Println(pkgs)
}

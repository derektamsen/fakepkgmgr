// This is a fake package manager that does nothing. It is only an excercise.
// It only prints what it would do using a static list.
package main

import "log"

// pkg defines the package name, version, and dependencies
type pkg struct {
	name         string
	version      string
	dependencies []pkg
}

func makeFakeList() *pkg {
	pkgs := pkg{
		name:    "golang",
		version: "2:1.12~1ubuntu1",
		dependencies: []pkg{
			{name: "golang-doc", version: "2:1.12~1ubuntu1"},
			{name: "golang-src", version: "2:1.12~1ubuntu1"},
			{
				name:         "golang-go",
				version:      "2:1.12~1ubuntu1",
				dependencies: []pkg{{name: "gcc", version: "4:9.2.1-3.1ubuntu1"}}},
		},
	}

	return &pkgs
}

func main() {
	pkgs := makeFakeList()

	log.Println(pkgs)
}

// This is a fake package manager that does nothing. It is only an excercise.
// It only prints what it would do using a static list.
package main

import (
	"fmt"
	"log"
)

// Pkg defines the package name, version, and dependencies
type Pkg struct {
	name         string
	version      string
	dependencies []Pkg
}

// installList is a list of full package names
type pkgList []string

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

// fullName returns the name + version of the package
func (p *Pkg) fullName() string {
	return fmt.Sprintf("%v_%v", p.name, p.version)
}

// install fake installs packages from installList
func install(l pkgList) error {
	for _, pkg := range l {
		log.Printf("Installing %s", pkg)
	}
	return nil
}

// getList returns a list of packages given Pkg
func getList(p *Pkg) pkgList {
	installPkgs := pkgList{p.fullName()}

	for _, pkg := range p.dependencies {
		installPkgs = append(installPkgs, pkg.fullName())
	}

	return installPkgs
}

func main() {
	pkgs := MakeFakeList()

	installPkgs := getList(pkgs)

	err := install(installPkgs)
	if err != nil {
		log.Panicln(err)
	}
}

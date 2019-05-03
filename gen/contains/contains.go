package main

import (
	"flag"
	"io"
	"text/template"

	"github.com/micvbang/go-helpy/gen"
)

type templateData struct {
	Type        string
	TypeName    string
	PackageName string
}

func main() {
	var d templateData
	flag.StringVar(&d.Type, "type", "", "The subtype used for the queue being generated")
	flag.StringVar(&d.PackageName, "package-name", "", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.Parse()

	t := template.Must(template.New("template").Parse(containsTemplate))

	gen.GenerateToPackage(d.PackageName, "gen_contains.go", func(w io.Writer) error {
		return t.Execute(w, d)
	})
}

const containsTemplate = `
package {{.PackageName}}

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []{{.Type}}, v {{.Type}}) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
`

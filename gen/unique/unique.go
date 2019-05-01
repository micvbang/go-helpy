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

	t := template.Must(template.New("unique").Parse(uniqueTemplate))

	gen.GenerateToPackage(d.PackageName, "gen_unique.go", func(w io.Writer) error {
		return t.Execute(w, d)
	})
}

const uniqueTemplate = `
package {{.PackageName}}

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique {{.Type}}s from the given input
func Unique(vs []{{.Type}}) []{{.Type}} {
	output := make([]{{.Type}}, 0, len(vs))
	seen := make(map[{{.Type}}]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
`

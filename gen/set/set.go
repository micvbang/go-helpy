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

	t := template.Must(template.New("template").Parse(setTemplate))

	gen.GenerateToPackage(d.PackageName, "gen_set.go", func(w io.Writer) error {
		return t.Execute(w, d)
	})
}

const setTemplate = `
package {{.PackageName}}

// Code generated. DO NOT EDIT.

type Set map[{{.Type}}]struct{}

// Contains returns true if Set contains v.
func (s Set) Contains(v {{.Type}}) bool {
	_, exists := s[v]
	return exists
}

// Intersect returns a Set of the intersection between two Sets.
func (s Set) Intersect(s2 Set) Set {
	short, long := s, s2
	if len(s2) < len(s) {
		short, long = s2, s
	}

	m := make(map[{{.Type}}]struct{}, len(short))
	for k := range short {
		if _, exists := long[k]; exists {
			m[k] = struct{}{}
		}
	}
	return m
}

// ToSet returns a lookup map for {{.Type}}.
func ToSet(vs []{{.Type}}) Set {
	m := make(map[{{.Type}}]struct{})
	for _, v := range vs {
		m[v] = struct{}{}
	}

	return m
}
`

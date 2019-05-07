package main

import (
	"flag"
	"io"
	"text/template"

	"github.com/micvbang/go-helpy/gen"
)

type templateData struct {
	Type        string
	PackageName string
}

func main() {
	var d templateData
	flag.StringVar(&d.Type, "type", "", "Name of the type")
	flag.StringVar(&d.PackageName, "package-name", "", "Name of the package to place the generated file in")
	flag.Parse()

	t := template.Must(template.New("minmax").Parse(minMaxTemplate))

	gen.GenerateToPackage(d.PackageName, "gen_minmax.go", func(w io.Writer) error {
		return t.Execute(w, d)
	})
}

const minMaxTemplate = `
package {{.PackageName}}

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v {{.Type}}, vs ...{{.Type}}) {{.Type}} {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}


// Max returns the maximum value from v and vs.
func Max(v {{.Type}}, vs ...{{.Type}}) {{.Type}} {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}
`

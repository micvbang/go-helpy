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

	abs := template.Must(template.New("abs").Parse(absTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_abs.go", func(w io.Writer) error {
		return abs.Execute(w, d)
	})

	absTest := template.Must(template.New("abs").Parse(absTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_abs_test.go", func(w io.Writer) error {
		return absTest.Execute(w, d)
	})
}

const absTemplate = `
package {{.PackageName}}

// Code generated. DO NOT EDIT.

import "github.com/micvbang/go-helpy"

// Abs returns the absolute value of v.
// NOTE: this method is deprecated. Use helpy.Abs instead.
func Abs(v {{.Type}}) {{.Type}} {
	return helpy.Abs(v)
}
`

const absTemplateTest = `
package {{.PackageName}}_test

import (
	"testing"
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected int
	}{
		"positive": {input: 42, expected: 42},
		"negative": {input: -123, expected: 123},
		"zero":     {input: 0, expected: 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := helpy.Abs(test.input)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
`

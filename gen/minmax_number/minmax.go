package main

import (
	"flag"
	"io"
	"strings"
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

	minMax := template.Must(template.New("minmax").Parse(minMaxTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_minmax.go", func(w io.Writer) error {
		return minMax.Execute(w, d)
	})

	funcMap := template.FuncMap{
		"HasPrefix": strings.HasPrefix,
	}
	minMaxTest := template.Must(template.New("minmax").Funcs(funcMap).Parse(minMaxTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_minmax_test.go", func(w io.Writer) error {
		return minMaxTest.Execute(w, d)
	})
}

const minMaxTemplate = `
package {{.PackageName}}

import "github.com/micvbang/go-helpy"

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
// NOTE: this method is deprecated. Use helpy.Min instead.
func Min(v {{.Type}}, vs ...{{.Type}}) {{.Type}} {
	return helpy.Min(v, vs...)
}


// Max returns the maximum value from v and vs.
// NOTE: this method is deprecated. Use helpy.Max instead.
func Max(v {{.Type}}, vs ...{{.Type}}) {{.Type}} {
	return helpy.Max(v, vs...)
}
`

const minMaxTemplateTest = `
package {{.PackageName}}_test

import (
	"testing"

	"github.com/micvbang/go-helpy/{{.PackageName}}"
)

// Code generated. DO NOT EDIT.

func TestMin(t *testing.T) {
	tests := map[string]struct {
		input []{{.Type}}
		expected {{.Type}}
	}{
		"in order": {input: []{{.Type}}{1, 2}, expected: 1},
		"reverse order": {input: []{{.Type}}{5, 2}, expected: 2},
		"3 arguments": {input: []{{.Type}}{10, 5, 15}, expected: 5},
		"4 arguments": {input: []{{.Type}}{1, 122, 100}, expected: 1},
		"10 arguments": {input: []{{.Type}}{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 0},
		{{if (HasPrefix .Type "int") }}
		"one negative": {input: []{{.Type}}{-5, 10}, expected: -5},
		"both negative": {input: []{{.Type}}{-5, -100}, expected: -100},
		"10 arguments negative": {input: []{{.Type}}{17, 25, -1, 0, -101, 127, 42, -13, 37, 69}, expected: -101},
		{{end}}
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := {{.PackageName}}.Min(v, vs...)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
    }	
}

func TestMax(t *testing.T) {
	tests := map[string]struct {
		input []{{.Type}}
		expected {{.Type}}
	}{
		"in order": {input: []{{.Type}}{1, 2}, expected: 2},
		"reverse order": {input: []{{.Type}}{5, 2}, expected: 5},
		"3 arguments": {input: []{{.Type}}{10, 5, 15}, expected: 15},
		"4 arguments": {input: []{{.Type}}{1, 122, 100}, expected: 122},
		"10 arguments": {input: []{{.Type}}{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 125},
		{{if (HasPrefix .Type "int") }}
		"one negative": {input: []{{.Type}}{-5, 10}, expected: 10},
		"both negative": {input: []{{.Type}}{-5, -100}, expected: -5},
		"10 arguments negative": {input: []{{.Type}}{17, 25, -1, 0, -101, 127, 42, -13, 37, 69}, expected: 127},
		{{end}}
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := {{.PackageName}}.Max(v, vs...)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
    }	
}
`

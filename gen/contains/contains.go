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

	contains := template.Must(template.New("template").Parse(containsTemplate))

	gen.GenerateToPackage(d.PackageName, "gen_contains.go", func(w io.Writer) error {
		return contains.Execute(w, d)
	})

	containsTest := template.Must(template.New("abs").Parse(containsTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_contains_test.go", func(w io.Writer) error {
		return containsTest.Execute(w, d)
	})
}

const containsTemplate = `
package {{.PackageName}}

// Code generated. DO NOT EDIT.

import "github.com/micvbang/go-helpy/slicey"

// Contains returns true if vs contains v.
// NOTE: this function is deprecated. Use slicey.Contains instead.
func Contains(vs []{{.Type}}, v {{.Type}}) bool {
	return slicey.Contains(vs, v)
}
`

const containsTemplateTest = `
package {{.PackageName}}_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

// TestContains verifies that Contains returns true only when f returns true.
func TestContains(t *testing.T) {
	tests := map[string]struct {
		haystack []int
		needle   int
		expected bool
	}{
		"nil haystack": {
			haystack: nil,
			expected: false,
		},
		"empty haystack": {
			haystack: []int{},
			expected: false,
		},
		"one element exists": {
			haystack: []int{42},
			needle:   42,
			expected: true,
		},
		"one element, doesn't exist": {
			haystack: []int{1337},
			needle:   42,
			expected: false,
		},
		"multiple elements, exists first": {
			haystack: []int{42, 1337, -1500, 17, 71, 18821390293},
			needle:   42,
			expected: true,
		},
		"multiple elements, exists middle": {
			haystack: []int{1337, -1500, 42, 17, 71, 18821390293},
			needle:   42,
			expected: true,
		},
		"multiple elements, exists last": {
			haystack: []int{1337, -1500, 17, 71, 18821390293, 42},
			needle:   42,
			expected: true,
		},
		"multiple elements, doesn't exist": {
			haystack: []int{-42, 1337, -1500, 17, 71, 18821390293},
			needle:   42,
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Contains(test.haystack, test.needle)

			if got != test.expected {
				t.Errorf("expected exist %v, got %v", test.expected, got)
			}
		})
	}
}
`

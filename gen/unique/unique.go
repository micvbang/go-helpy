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

	uniqueTest := template.Must(template.New("unique").Parse(uniqueTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_unique_test.go", func(w io.Writer) error {
		return uniqueTest.Execute(w, d)
	})
}

const uniqueTemplate = `
package {{.PackageName}}

import (
	"github.com/micvbang/go-helpy/slicey"
)

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique {{.Type}}s from the given input
// NOTE: this function is deprecated. Use slicey.Unique instead.
func Unique(vs []{{.Type}}) []{{.Type}} {
	return slicey.Unique(vs)
}
`

const uniqueTemplateTest = `
package {{.PackageName}}_test

import (
	"testing"
	"fmt"

	"github.com/micvbang/go-helpy/slicey"
	"github.com/micvbang/go-helpy/{{.PackageName}}"
)

// Code generated. DO NOT EDIT.

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []{{.Type}}
		expected []{{.Type}}
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			got := {{.PackageName}}.Unique(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
    }	
}

func UniqueGenerateTestCase(n int) (test []{{.Type}}, expected []{{.Type}}) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]{{.Type}}, n)
	expected = make([]{{.Type}}, 0, n)
	seen := make(map[{{.Type}}]struct{})

	for i := 0; i < n; i++ {
		v := {{.PackageName}}.RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok{
			seen[v] = struct{}{}
			expected = append(expected,  v)
		}
	}

	return test, expected
}
`

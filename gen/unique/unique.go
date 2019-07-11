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

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique {{.Type}}s from the given input
func Unique(vs []{{.Type}}) []{{.Type}} {
	output := make([]{{.Type}}, 0, len(vs))
	seen := make(map[{{.Type}}]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			output = append(output, v)
		}
	}

	return output
}
`

const uniqueTemplateTest = `
package {{.PackageName}}

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func UniqueGenerateTestCase(n int) (test []{{.Type}}, expected []{{.Type}}) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]{{.Type}}, n)
	expected = make([]{{.Type}}, 0, n)
	seen := make(map[{{.Type}}]struct{})

	for i := 0; i < n; i++ {
		v := RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok{
			seen[v] = struct{}{}
			expected = append(expected,  v)
		}
	}

	return test, expected
}

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

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			require.Equal(t, tc.expected, Unique(tc.input))
		})
    }	
}
`

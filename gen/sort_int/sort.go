package main

import (
	"flag"
	"io"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

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

	funcMap := template.FuncMap{
		"Title": cases.Title(language.English).String,
	}
	sort := template.Must(template.New("template").Funcs(funcMap).Parse(sortTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_sort.go", func(w io.Writer) error {
		return sort.Execute(w, d)
	})

	sortTest := template.Must(template.New("template").Parse(sortTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_sort_test.go", func(w io.Writer) error {
		return sortTest.Execute(w, d)
	})
}

const sortTemplate = `
package {{.PackageName}}

import (
	"sort"
)

// Code generated. DO NOT EDIT.

type {{.Type | Title}}Slice []{{.Type}}

func (p {{.Type | Title}}Slice) Len() int           { return len(p) }
func (p {{.Type | Title}}Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p {{.Type | Title}}Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort sorts data in place. Not guaranteed to be stable.
func Sort(vs []{{.Type}}) []{{.Type}} {
	sort.Sort({{.Type | Title}}Slice(vs))
	return vs
}
`

const sortTemplateTest = `
package {{.PackageName}}

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestSort(t *testing.T) {
	tests := map[string]struct {
		input []{{.Type}}
		expected []{{.Type}}
	}{
		"in order": { 
			input: []{{.Type}}{1, 2, 3, 4, 5},
			expected: []{{.Type}}{1, 2, 3, 4, 5},
		},
		"reverse order": { 
			input: []{{.Type}}{5, 4, 3, 2, 1},
			expected: []{{.Type}}{1, 2, 3, 4, 5},
		},
		"random order": { 
			input: []{{.Type}}{4, 1, 2, 5, 3},
			expected: []{{.Type}}{1, 2, 3, 4, 5},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, Sort(tc.input))
		})
    }
}
`

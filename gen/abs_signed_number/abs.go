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

	abs := template.Must(template.New("abs").Parse(absTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_abs.go", func(w io.Writer) error {
		return abs.Execute(w, d)
	})

	funcMap := template.FuncMap{
		"HasPrefix": strings.HasPrefix,
	}
	absTest := template.Must(template.New("abs").Funcs(funcMap).Parse(absTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_abs_test.go", func(w io.Writer) error {
		return absTest.Execute(w, d)
	})
}

const absTemplate = `
package {{.PackageName}}

// Code generated. DO NOT EDIT.

// Abs returns the absolute value of v.
func Abs(v {{.Type}}) {{.Type}} {
	if v < 0 {
		return -v
	}
	return v
}
`

const absTemplateTest = `
package {{.PackageName}}

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		input {{.Type}}
		expected {{.Type}}
	}{
		"positive": {input: 42, expected: 42},
		"negative": {input: -123, expected: 123},
		"zero": {input: 0, expected: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, Abs(tc.input))
		})
    }	
}
`

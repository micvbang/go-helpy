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

const minMaxTemplateTest = `
package {{.PackageName}}

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		"4 arguments": {input: []{{.Type}}{1, 500, 100}, expected: 1},
		"10 arguments": {input: []{{.Type}}{17, 25, 1, 0, 101, 1337, 42, 13, 37, 69}, expected: 0},
		{{if (HasPrefix .Type "int") }}
		"one negative": {input: []{{.Type}}{-5, 10}, expected: -5},
		"both negative": {input: []{{.Type}}{-5, -100}, expected: -100},
		"10 arguments negative": {input: []{{.Type}}{17, 25, -1, 0, -101, 1337, 42, -13, 37, 69}, expected: -101},
		{{end}}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := tc.input[0], tc.input[1:len(tc.input)]
			require.Equal(t, tc.expected, Min(v, vs...))
		})
    }	
}
`

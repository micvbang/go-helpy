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

	random := template.Must(template.New("template").Parse(randomTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_random.go", func(w io.Writer) error {
		return random.Execute(w, d)
	})

	funcMap := template.FuncMap{
		"HasPrefix": strings.HasPrefix,
	}
	randomTest := template.Must(template.New("template").Funcs(funcMap).Parse(randomTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_random_test.go", func(w io.Writer) error {
		return randomTest.Execute(w, d)
	})
}

const randomTemplate = `
package {{.PackageName}}

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type {{.Type}}
func Random() {{.Type}} {
	return {{.Type}}(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type {{.Type}}. It panics if n <= 0.
func RandomN(n {{.Type}}) {{.Type}} {
	return {{.Type}}(rand.Int63n(int64(n)))
}
`

const randomTemplateTest = `
package {{.PackageName}}

import (
	"testing"
)

// Code generated. DO NOT EDIT.

// TestRandom simply runs Random to ensure that it compiles during testing.
func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		Random()
	}
}

// TestRandomN runs RandomN and ensures that output is within expected range.
func TestRandomN(t *testing.T) {
	for i := 1; i < 127; i++ {
		v := {{.Type}}(i)
		got := RandomN(v) 
		if got >= v {
			t.Errorf("expected value < %v, got %v", v, got)
		}
		{{if (HasPrefix .Type "int") }}
			if got >= v {
				t.Errorf("expected value > 0, got %v", got)
			}
		{{end}}
	}
}
`

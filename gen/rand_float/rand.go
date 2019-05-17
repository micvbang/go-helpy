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

	random := template.Must(template.New("template").Parse(randomTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_random.go", func(w io.Writer) error {
		return random.Execute(w, d)
	})

	randomTest := template.Must(template.New("template").Parse(randomTemplateTest))
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

// Random returns a random value of type {{.Type}}
func Random() {{.Type}} {
	return {{.Type}}(rand.Float64())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type {{.Type}}.
func RandomN(n {{.Type}}) {{.Type}} {
	return {{.Type}}(rand.Float64()) * n
}
`

const randomTemplateTest = `
package {{.PackageName}}

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	for i := 1; i < 10000; i++ {
		n := RandomN({{.Type}}(i)) 
		require.True(t, 0 <= n && n < {{.Type}}(i))
	}
}
`

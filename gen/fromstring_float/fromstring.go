package main

import (
	"flag"
	"io"
	"log"
	"text/template"

	"github.com/micvbang/go-helpy/gen"
	"github.com/micvbang/go-helpy/slicey"
)

type templateData struct {
	Type        string
	PackageName string
	BitSize     int
}

func main() {
	var d templateData
	flag.StringVar(&d.Type, "type", "", "Name of the type")
	flag.StringVar(&d.PackageName, "package-name", "", "Name of the package to place the generated file in")
	flag.IntVar(&d.BitSize, "bit-size", 0, "Number of bits used to represent number")
	flag.Parse()

	validBitSizes := []int{32, 64}
	if !slicey.Contains(validBitSizes, d.BitSize) {
		log.Fatalf("Bit size was '%d', but must be one of %v", d.BitSize, validBitSizes)
	}

	fromstring := template.Must(template.New("fromstring").Parse(fromStringTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_fromstring.go", func(w io.Writer) error {
		return fromstring.Execute(w, d)
	})

	fromstringTest := template.Must(template.New("fromstring").Parse(absTemplateTest))
	gen.GenerateToPackage(d.PackageName, "gen_fromstring_test.go", func(w io.Writer) error {
		return fromstringTest.Execute(w, d)
	})
}

const fromStringTemplate = `
package {{.PackageName}}

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a {{.Type}}.
func FromString(s string) ({{.Type}}, error) {
	v, err := strconv.ParseFloat(s, {{.BitSize}})
	return {{.Type}}(v), err
}

// FromStringOrDefault parses s and returns a {{.Type}}.
func FromStringOrDefault(s string, defaultVal {{.Type}}) {{.Type}} {
	v, err := strconv.ParseFloat(s, {{.BitSize}})
	if err != nil {
		return defaultVal
	}
	return {{.Type}}(v)
}
`

const absTemplateTest = `
package {{.PackageName}}_test

import (
	"testing"
	
	"github.com/micvbang/go-helpy/{{.PackageName}}"
)

// Code generated. DO NOT EDIT.

func TestFromString(t *testing.T) {
	tests := map[string]struct {
		input string
		expected {{.Type}}
	}{
		"positive": {input: "42.1337", expected: 42.1337},
		"zero": {input: "0", expected: 0},
		"negative": {input: "-123.321", expected: -123.321},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := {{.PackageName}}.FromString(test.input)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if v != test.expected {
				t.Errorf("expected %v, got %v", test.expected, v)
			}
		})
    }	
}

func TestFromStringOrDefault(t *testing.T) {
	tests := map[string]struct {
		input string
		defaultVal {{.Type}}
		expected {{.Type}}
	}{
		"positive": {input: "42.1337", expected: 42.1337},
		"zero": {input: "0", expected: 0},
		"default": {input: "sdf", defaultVal: 42.5, expected: 42.5},
		"negative": {input: "-123.321", expected: -123.321},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := {{.PackageName}}.FromStringOrDefault(test.input, test.defaultVal)
			if v != test.expected {
				t.Errorf("expected %v, got %v", test.expected, v)
			}
		})
    }	
}
`

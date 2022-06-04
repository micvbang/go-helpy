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
	TypeName    string
	PackageName string
}

func main() {
	var d templateData
	flag.StringVar(&d.Type, "type", "", "The subtype used for the queue being generated")
	flag.StringVar(&d.PackageName, "package-name", "", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.Parse()

	set := template.Must(template.New("template").Parse(setTemplate))
	gen.GenerateToPackage(d.PackageName, "gen_set.go", func(w io.Writer) error {
		return set.Execute(w, d)
	})

	if strings.HasPrefix(d.Type, "int") {
		setTest := template.Must(template.New("template").Parse(setTemplateTest))
		gen.GenerateToPackage(d.PackageName, "gen_set_test.go", func(w io.Writer) error {
			return setTest.Execute(w, d)
		})
	}
}

const setTemplate = `
package {{.PackageName}}

import "github.com/micvbang/go-helpy"

// Code generated. DO NOT EDIT.

// NOTE: this type is deprecated. Use helpy.Set instead
type Set helpy.Set[{{.Type}}]

func (s Set) t() helpy.Set[{{.Type}}] {
	return helpy.Set[{{.Type}}](s)
}

func (s Set) Intersect(s2 Set) Set {
	return Set(s.t().Intersect(s2.t()))
}

func (s Set) Union(s2 Set) Set {
	return Set(s.t().Union(s2.t()))
}

func (s Set) Contains(v {{.Type}}) bool {
	return s.t().Contains(v)
}

func (s Set) Equal(s2 Set) bool {
	return s.t().Equal(s2.t())
}

// ToSet returns a lookup map for {{.Type}}.
// NOTE: this function is deprecated. Use helpy.ToSet instead.
func ToSet(vs []{{.Type}}) Set {
	return Set(helpy.ToSet(vs))
}

// MakeSet returns a lookup map for {{.Type}}
// NOTE: this function is deprecated. Use helpy.MakeSet instead.
func MakeSet(vs ...{{.Type}}) Set {
	return Set(helpy.ToSet(vs))
}
`

const setTemplateTest = `
package {{.PackageName}}

import (
	"testing"
)

// Code generated. DO NOT EDIT.


func TestSetIntersect(t *testing.T) {
	tests := map[string]struct {
		s1 Set
		s2 Set
		expected Set
	}{
		"none common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{5, 6, 7, 8}),
			expected: ToSet([]{{.Type}}{}),
		},
		"one common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{2, 6, 7, 8}),
			expected: ToSet([]{{.Type}}{2}),
		},
		"multiple common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{2, 1, 7, 4}),
			expected: ToSet([]{{.Type}}{1, 2, 4}),
		},
		"all common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{4, 3, 2, 1}),
			expected: ToSet([]{{.Type}}{1, 2, 3, 4}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(test.s1)
			s2Len := len(test.s2)

			got := test.s1.Intersect(test.s2)
			if !got.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}

			// Ensure original sets weren't modified
			if len(test.s1) != s1Len {
				t.Errorf("expected set to be of size %v, got %v", s1Len, test.s1)
			}

			if len(test.s2) != s2Len {
				t.Errorf("expected set to be of size %v, got %v", s2Len, test.s2)
			}
		})
    }	
}


func TestSetUnion(t *testing.T) {
	tests := map[string]struct {
		s1 Set
		s2 Set
		expected Set
	}{
		"none common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{5, 6, 7, 8}),
			expected: ToSet([]{{.Type}}{1, 2, 3, 4, 5, 6, 7, 8}),
		},
		"one common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{2, 6, 7, 8}),
			expected: ToSet([]{{.Type}}{1, 2, 3, 4, 6, 7, 8}),
		},
		"multiple common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{2, 1, 7, 4}),
			expected: ToSet([]{{.Type}}{1, 2, 3, 4, 7}),
		},
		"all common": {
			s1: ToSet([]{{.Type}}{1, 2, 3, 4}),
			s2: ToSet([]{{.Type}}{4, 3, 2, 1}),
			expected: ToSet([]{{.Type}}{1, 2, 3, 4}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(test.s1)
			s2Len := len(test.s2)

			got := test.s1.Union(test.s2)
			if !got.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}

			// Ensure original sets weren't modified
			if len(test.s1) != s1Len {
				t.Errorf("expected set to be of size %v, got %v", s1Len, test.s1)
			}

			if len(test.s2) != s2Len {
				t.Errorf("expected set to be of size %v, got %v", s2Len, test.s2)
			}
		})
    }	
}
`

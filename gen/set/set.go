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

// Code generated. DO NOT EDIT.

type Set map[{{.Type}}]struct{}

// Contains returns true if Set contains v.
func (s Set) Contains(v {{.Type}}) bool {
	_, exists := s[v]
	return exists
}

// Intersect returns the intersection of two Sets.
func (s Set) Intersect(s2 Set) Set {
	short, long := s, s2
	if len(s2) < len(s) {
		short, long = s2, s
	}

	m := make(map[{{.Type}}]struct{}, len(short))
	for k := range short {
		if _, exists := long[k]; exists {
			m[k] = struct{}{}
		}
	}
	return m
}

// Union returns the union of two Sets.
func (s Set) Union(s2 Set) Set {
	m := make(map[{{.Type}}]struct{}, len(s) + len(s2))
	for k := range s {
		m[k] = struct{}{}
	}

	for k := range s2 {
		m[k] = struct{}{}
	}

	return m
}

// Equal checks whether all 
func (s Set) Equal(s2 Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for v := range s {
		if _, ok := s2[v]; !ok {
			return false
		}
	}

	return true
}

// ToSet returns a lookup map for {{.Type}}.
func ToSet(vs []{{.Type}}) Set {
	m := make(map[{{.Type}}]struct{})
	for _, v := range vs {
		m[v] = struct{}{}
	}

	return m
}
`

const setTemplateTest = `
package {{.PackageName}}

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(tc.s1)
			s2Len := len(tc.s2)

			require.True(t, tc.s1.Intersect(tc.s2).Equal(tc.expected))

			// Ensure original sets weren't modified
			require.Equal(t, len(tc.s1), s1Len)
			require.Equal(t, len(tc.s2), s2Len)
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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(tc.s1)
			s2Len := len(tc.s2)

			require.True(t, tc.s1.Union(tc.s2).Equal(tc.expected))

			// Ensure original sets weren't modified
			require.Equal(t, len(tc.s1), s1Len)
			require.Equal(t, len(tc.s2), s2Len)
		})
    }	
}
`

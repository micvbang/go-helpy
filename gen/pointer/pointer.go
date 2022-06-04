package main

import (
	"flag"
	"io"
	"text/template"
	"unicode"

	"github.com/micvbang/go-helpy/gen"
	"github.com/micvbang/go-helpy/internal/stringsy"
)

type templateData struct {
	Type        string
	TypeName    string
	PackageName string
	Import      string
}

func main() {
	var d templateData
	flag.StringVar(&d.Type, "type", "", "Name of the type")
	flag.StringVar(&d.TypeName, "type-name", "", "User friendly name of the type")
	flag.StringVar(&d.PackageName, "package-name", "", "Name of the package to place the generated file in")
	flag.StringVar(&d.Import, "import", "", "Name of package to import")
	flag.Parse()

	funcMap := template.FuncMap{
		"Title": func(s string) string {
			return stringsy.TitleCasing(s, unicode.ToUpper)
		},
	}

	t := template.Must(template.New("pointer").Funcs(funcMap).Parse(pointerTemplate))

	gen.GenerateToPackage(d.PackageName, "gen_pointer.go", func(w io.Writer) error {
		return t.Execute(w, d)
	})
}

const pointerTemplate = `
package {{.PackageName}}

import (
	"github.com/micvbang/go-helpy"
	{{if .Import}}
		"{{.Import}}"
	{{end}}

)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given {{.Type}}.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v {{.Type}}) *{{.Type}} {
	return helpy.Pointer(v)
}

// {{if .TypeName}}{{.TypeName}}{{else}}{{.Type | Title}}{{end}}OrDefault returns {{.Type}} if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func {{if .TypeName}}{{.TypeName}}{{else}}{{.Type | Title}}{{end}}OrDefault(v *{{.Type}}, defaultVal {{.Type}}) *{{.Type}} {
	return helpy.DerefOrValue(v, defaultVal)
}
`

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
		"Title": strings.Title,
	}

	t := template.Must(template.New("pointer").Funcs(funcMap).Parse(pointerTemplate))

	gen.GenerateToPackage(d.PackageName, "gen_pointer.go", func(w io.Writer) error {
		return t.Execute(w, d)
	})
}

const pointerTemplate = `
package {{.PackageName}}

{{if .Import}}
import (
	"{{.Import}}"
)
{{end}}

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given {{.Type}}.
func Pointer(v {{.Type}}) *{{.Type}} {
	return &v
}

// {{if .TypeName}}{{.TypeName}}{{else}}{{.Type | Title}}{{end}} dereferences and returns {{.Type}}. The {{.Type}} default value is returned if v is nil.
func {{if .TypeName}}{{.TypeName}}{{else}}{{.Type | Title}}{{end}}(v *{{.Type}}) {{.Type}} {
	if v == nil {
		var dv {{.Type}}
		return dv
	}
	return *v
}

// {{if .TypeName}}{{.TypeName}}{{else}}{{.Type | Title}}{{end}}OrDefault dereferences and returns {{.Type}}. defaultVal is returned if v is nil.
func {{if .TypeName}}{{.TypeName}}{{else}}{{.Type | Title}}{{end}}OrDefault(v *{{.Type}}, defaultVal {{.Type}}) {{.Type}} {
	if v == nil {
		var dv {{.Type}}
		return dv
	}
	return *v
}
`

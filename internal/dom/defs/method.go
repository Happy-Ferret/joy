package defs

import (
	"strings"

	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
	"github.com/matthewmueller/golly/internal/gen"
	"github.com/pkg/errors"
)

var _ Method = (*method)(nil)

// NewMethod creates a new method
func NewMethod(index index.Index, data *raw.Method, receiver Interface) Method {
	return &method{
		index: index,
		data:  data,
		recv:  receiver,
	}
}

// Method interface
type Method interface {
	def.Definition
}

// method struct
type method struct {
	data *raw.Method
	pkg  string
	file string

	index   index.Index
	comment string
	recv    Interface
}

func (d *method) ID() string {
	return d.recv.ID() + " " + d.data.Name
}

func (d *method) Name() string {
	return d.data.Name
}

func (d *method) Kind() string {
	return "METHOD"
}

func (d *method) Type() (string, error) {
	return d.index.Coerce(d.data.Name)
}

func (d *method) SetPackage(pkg string) {
	d.pkg = pkg
}
func (d *method) GetPackage() string {
	return d.pkg
}

func (d *method) SetFile(file string) {
	d.file = file
}
func (d *method) GetFile() string {
	return d.file
}

// Dependencies fn
func (d *method) Dependencies() (defs []def.Definition, e error) {
	for _, param := range d.data.Params {
		if def := d.index.Find(param.Type); def != nil {
			defs = append(defs, def)
		}
	}
	if def := d.index.Find(d.data.Type); def != nil {
		defs = append(defs, def)
	}
	return defs, nil
}

// Generate fn
func (d *method) Generate() (string, error) {
	data := struct {
		Recv   string
		Name   string
		Params []gen.Vartype
		Result gen.Vartype
	}{
		Recv: gen.Pointer(d.recv.Name()),
		Name: gen.Capitalize(d.data.Name),
	}

	for _, param := range d.data.Params {
		t, err := d.index.Coerce(param.Type)
		if err != nil {
			return "", errors.Wrapf(err, "error coercing param")
		}
		data.Params = append(data.Params, gen.Vartype{
			Var:      gen.Identifier(param.Name),
			Optional: param.Optional,
			Type:     t,
		})
	}

	t, e := d.index.Coerce(d.data.Type)
	if e != nil {
		return "", e
	}
	data.Result = gen.Vartype{
		Var:  gen.Variable(t),
		Type: t,
	}

	async := strings.Contains(d.data.Type, "Promise<")
	if t == "" {
		if async {
			return gen.Generate("method/"+d.data.Name, data, `
				func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
					js.Rewrite("await $<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
				}
			`)
		}

		return gen.Generate("method/"+d.data.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
				js.Rewrite("$<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
			}
		`)
	}

	if async {
		return gen.Generate("method/"+d.data.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ vt .Result }}) {
				js.Rewrite("await $<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
				return {{ .Result.Var }}
			}
		`)
	}
	return gen.Generate("method/"+d.data.Name, data, `
		func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ vt .Result }}) {
			js.Rewrite("$<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
			return {{ .Result.Var }}
		}
	`)
}
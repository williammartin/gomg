package generator

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/williammartin/jsonschema"
)

//go:generate counterfeiter . Reflector

type Reflector interface {
	Reflect(v interface{}) *jsonschema.Schema
}

type SchemaGenerator struct {
	Reflector Reflector
}

func (g *SchemaGenerator) Generate(v interface{}) (io.Reader, error) {
	schema := g.Reflector.Reflect(v)

	b, err := json.Marshal(schema)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

package generator

import (
	"bytes"
	"encoding/json"
	"io"
)

type DocumentGenerator struct{}

func (g *DocumentGenerator) Generate(v interface{}) (io.Reader, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

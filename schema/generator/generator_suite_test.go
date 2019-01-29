package generator_test

import (
	"io"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generator Suite")
}

func readAll(reader io.Reader) string {
	b, err := ioutil.ReadAll(reader)
	Expect(err).NotTo(HaveOccurred())

	return string(b)
}

type Foo struct {
	Bar bool `json:"bar,omitempty"`
}

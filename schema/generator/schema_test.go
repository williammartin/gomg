package generator_test

import (
	"io"
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/williammartin/gomg/schema/generator"
	"github.com/williammartin/gomg/schema/generator/generatorfakes"
	"github.com/williammartin/jsonschema"
)

var _ = Describe("Schema Generator", func() {

	var (
		fakeReflector *generatorfakes.FakeReflector

		schemaGenerator *SchemaGenerator
		value           interface{}

		schema io.Reader
		err    error
	)

	BeforeEach(func() {
		fakeReflector = new(generatorfakes.FakeReflector)
		fakeReflector.ReflectReturns(&jsonschema.Schema{Type: &jsonschema.Type{Pattern: "test"}})
		schemaGenerator = &SchemaGenerator{
			Reflector: fakeReflector,
		}

		value = &Foo{Bar: true}
	})

	JustBeforeEach(func() {
		schema, err = schemaGenerator.Generate(value)
	})

	It("doesn't return an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("reflects the provided structure", func() {
		Expect(fakeReflector.ReflectCallCount()).NotTo(BeZero())
		Expect(fakeReflector.ReflectArgsForCall(0)).To(Equal(value))
	})

	It("marshals the reflected schema to json and returns a reader", func() {
		Expect(readAll(schema)).To(Equal(`{"pattern":"test"}`))
	})

	When("marshaling returns an error", func() {
		BeforeEach(func() {
			fakeReflector.ReflectReturns(&jsonschema.Schema{Type: &jsonschema.Type{Default: math.Inf(1)}})
		})

		It("propagates the error", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})

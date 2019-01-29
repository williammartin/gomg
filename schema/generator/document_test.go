package generator_test

import (
	"io"
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/williammartin/gomg/schema/generator"
)

var _ = Describe("Document Generator", func() {
	var (
		documentGenerator *DocumentGenerator
		value             interface{}

		document io.Reader
		err      error
	)

	BeforeEach(func() {
		documentGenerator = &DocumentGenerator{}
		value = &Foo{Bar: true}
	})

	JustBeforeEach(func() {
		document, err = documentGenerator.Generate(value)
	})

	It("doesn't return an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("marshals the reflected schema to json and returns a reader", func() {
		Expect(readAll(document)).To(Equal(`{"bar":true}`))
	})

	When("marshaling returns an error", func() {
		BeforeEach(func() {
			value = math.Inf(1)
		})

		It("propagates the error", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})

package schema_test

import (
	"errors"
	"io"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/williammartin/gomg/schema"
)

var _ = Describe("Validator", func() {
	var (
		validator *Validator

		schemaReader   io.Reader
		documentReader io.Reader

		result *Result
		err    error
	)

	BeforeEach(func() {
		validator = &Validator{}
	})

	JustBeforeEach(func() {
		result, err = validator.Validate(schemaReader, documentReader)
	})

	BeforeEach(func() {
		schemaReader = strings.NewReader(generateSchema())
		documentReader = strings.NewReader(`{"bar":true}`)
	})

	When("passed a valid schema and document", func() {
		It("doesn't return an error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns as valid", func() {
			Expect(result.IsValid).To(BeTrue())
		})

		It("returns no validation errors", func() {
			Expect(result.Errors).To(HaveLen(0))
		})
	})

	When("passed a document that doesn't pass schema validation", func() {
		BeforeEach(func() {
			documentReader = strings.NewReader(`{}`)
		})

		It("returns as not valid", func() {
			Expect(result.IsValid).To(BeFalse())
		})

		It("returns validation errors", func() {
			Expect(result.Errors).To(ConsistOf("(root): bar is required"))
		})
	})

	When("passed an schema that isn't json", func() {
		BeforeEach(func() {
			schemaReader = strings.NewReader("not-a-valid-schema")
		})

		It("returns an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	When("passed a document that isn't json", func() {
		BeforeEach(func() {
			documentReader = strings.NewReader("not-a-valid-document")
		})

		It("returns an error", func() {
			Expect(err).To(HaveOccurred())
		})

	})

	When("passed a schema reader that errors", func() {
		BeforeEach(func() {
			schemaReader = &ErrorReader{}
		})

		It("propagates the error", func() {
			Expect(err).To(MatchError("boom"))
		})
	})

	When("passed a document reader that errors", func() {
		BeforeEach(func() {
			documentReader = &ErrorReader{}
		})

		It("propagates the error", func() {
			Expect(err).To(MatchError("boom"))
		})
	})
})

func generateSchema() string {
	return `{"$schema":"http://json-schema.org/draft-07/schema#","$ref":"#/definitions/schema_test.Foo","definitions":{"schema_test.Foo":{"required":["bar"],"properties":{"bar":{"type":"boolean"}},"additionalProperties":false,"type":"object"}}}`
}

type ErrorReader struct{}

func (*ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("boom")
}

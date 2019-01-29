package validate_test

import (
	"errors"
	"io"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/williammartin/gomg/schema"
	. "github.com/williammartin/gomg/validate"
	"github.com/williammartin/gomg/validate/validatefakes"
	"github.com/williammartin/omg"
)

var _ = Describe("Actor", func() {

	var (
		fakeSchemaGenerator   *validatefakes.FakeSchemaGenerator
		fakeDocumentGenerator *validatefakes.FakeDocumentGenerator
		fakeValidator         *validatefakes.FakeSchemaValidator

		actor *Actor
	)

	BeforeEach(func() {
		fakeSchemaGenerator = new(validatefakes.FakeSchemaGenerator)
		fakeDocumentGenerator = new(validatefakes.FakeDocumentGenerator)
		fakeValidator = new(validatefakes.FakeSchemaValidator)
		actor = &Actor{
			SchemaGenerator:   fakeSchemaGenerator,
			DocumentGenerator: fakeDocumentGenerator,
			SchemaValidator:   fakeValidator,
		}
	})

	Describe("ValidateMicroservice", func() {
		var (
			validationResult *schema.Result
			err              error
		)

		JustBeforeEach(func() {
			validationResult, err = actor.ValidateMicroservice(&omg.Microservice{})
		})

		It("generates the schema", func() {
			Expect(fakeSchemaGenerator.GenerateCallCount()).NotTo(BeZero())
			Expect(fakeSchemaGenerator.GenerateArgsForCall(0)).To(Equal(&omg.Microservice{}))
		})

		When("schema generation succeeds", func() {
			var generatedSchema io.Reader

			BeforeEach(func() {
				generatedSchema = strings.NewReader("some-schema")
				fakeSchemaGenerator.GenerateReturns(generatedSchema, nil)
			})

			It("generates the document", func() {
				Expect(fakeDocumentGenerator.GenerateCallCount()).NotTo(BeZero())
				Expect(fakeDocumentGenerator.GenerateArgsForCall(0)).To(Equal(&omg.Microservice{}))
			})

			When("document generation succeeds", func() {
				var generatedDocument io.Reader

				BeforeEach(func() {
					generatedDocument = strings.NewReader("some-document")
					fakeDocumentGenerator.GenerateReturns(generatedDocument, nil)
				})

				It("calls the validator with the schema from the generator", func() {
					Expect(fakeValidator.ValidateCallCount()).NotTo(BeZero())
					actualSchema, actualDocument := fakeValidator.ValidateArgsForCall(0)
					Expect(actualSchema).To(Equal(generatedSchema))
					Expect(actualDocument).To(Equal(generatedDocument))
				})

				When("validation succeeds", func() {
					result := &schema.Result{
						IsValid: true,
						Errors:  schema.ValidationErrors{"one", "two", "three"},
					}

					BeforeEach(func() {
						fakeValidator.ValidateReturns(result, nil)
					})

					It("doesn't return an error", func() {
						Expect(err).NotTo(HaveOccurred())
					})

					It("returns the result", func() {
						Expect(validationResult.IsValid).To(BeTrue())
						Expect(validationResult.Errors).To(ConsistOf("one", "two", "three"))
					})
				})

				When("schema validation fails", func() {
					BeforeEach(func() {
						fakeValidator.ValidateReturns(nil, errors.New("boom"))
					})

					It("propagates the error", func() {
						Expect(err).To(MatchError("boom"))
					})
				})
			})

			When("document generation fails", func() {
				BeforeEach(func() {
					fakeDocumentGenerator.GenerateReturns(nil, errors.New("boom"))
				})

				It("propagates the error", func() {
					Expect(err).To(MatchError("boom"))
				})
			})
		})

		When("schema generation fails", func() {
			BeforeEach(func() {
				fakeSchemaGenerator.GenerateReturns(nil, errors.New("boom"))
			})

			It("propagates the error", func() {
				Expect(err).To(MatchError("boom"))
			})
		})

	})
})

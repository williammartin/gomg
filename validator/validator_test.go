package validator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/williammartin/gomg/validator"
	"github.com/williammartin/omg"
)

var _ = Describe("Validator", func() {

	var (
		validator    *Validator
		microservice *omg.Microservice

		result *Result
		err    error
	)

	BeforeEach(func() {
		validator = &Validator{
			MicroserviceSchema: "file:///Users/wmartin/go/src/github.com/williammartin/omg/schemas/microservice.json",
		}
	})

	JustBeforeEach(func() {
		result, err = validator.Validate(microservice)
	})

	When("passed a valid microservice", func() {
		BeforeEach(func() {
			microservice = generateValidMicroservice()
		})

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

	When("passed an invalid microservice", func() {
		BeforeEach(func() {
			microservice = &omg.Microservice{}
		})

		It("doesn't return an error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns as invalid", func() {
			Expect(result.IsValid).To(BeFalse())
		})

		It("returns validation errors", func() {
			Expect(len(result.Errors)).NotTo(BeZero())
		})
	})

	When("given a bad schema", func() {
		BeforeEach(func() {
			validator = &Validator{
				MicroserviceSchema: "",
			}
		})

		It("returns an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})

func generateValidMicroservice() *omg.Microservice {
	return &omg.Microservice{
		OMG: 1,
		Info: &omg.Info{
			Version:     "0.0.1",
			Title:       "Test Microservice",
			Description: "A Test Microservice",
			License: &omg.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
		},
	}
}

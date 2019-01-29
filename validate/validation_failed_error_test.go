package validate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/williammartin/gomg/validate"
)

var _ = Describe("ValidationFailedError", func() {
	It("produces useful error output from the validation errors", func() {
		validationErrors := []string{"first", "second"}
		err := &ValidationFailedError{ValidationErrors: validationErrors}

		Expect(err.Error()).To(ContainSubstring("validation errors occurred"))
		Expect(err.Error()).To(ContainSubstring(" - first"))
		Expect(err.Error()).To(ContainSubstring(" - second"))
	})
})

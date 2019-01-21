package ui_test

import (
	"bytes"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "github.com/williammartin/gomg/ui"
)

var _ = Describe("UI", func() {

	var testUI *UI

	BeforeEach(func() {
		testUI = &UI{
			Out: NewBuffer(),
			Err: NewBuffer(),
		}
	})

	Describe("DisplayText", func() {
		It("prints text with templated values to the out buffer", func() {
			testUI.DisplayText("This is a test for the {{.Struct}} struct", map[string]interface{}{"Struct": "UI"})
			Expect(testUI.Out).To(Say("This is a test for the UI struct\n"))
		})
	})

	Describe("DisplayError", func() {
		It("prints text to the err buffer", func() {
			testUI.DisplayError(errors.New("error text"))
			Expect(testUI.Err).To(Say("error text\n"))
		})
	})

	Describe("DisplayErrorAndFailed", func() {
		It("prints text to the err buffer and FAILED", func() {
			testUI.DisplayErrorAndFailed(errors.New("error text"))
			Expect(testUI.Err).To(Say("error text\n"))
			Expect(testUI.Err).To(Say("FAILED\n"))
		})
	})

	Describe("DisplaySuccess", func() {
		It("prints SUCCESS to the out buffer", func() {
			testUI.DisplaySuccess()
			Expect(testUI.Out).To(Say("SUCCESS\n"))
		})
	})

	Describe("DisplayFailed", func() {
		It("prints FAILED to the err buffer", func() {
			testUI.DisplayFailed()
			Expect(testUI.Err).To(Say("FAILED\n"))
		})
	})

	Describe("DisplayNewline", func() {
		It("prints a newline to the out buffer", func() {
			testUI.DisplayNewline()
			Expect(testUI.Out).To(Say("\n"))
		})
	})

	Describe("DisplayStream", func() {
		It("copies from the reader to the out buffer", func() {
			output := bytes.NewBuffer([]byte("test-output"))
			testUI.DisplayStream(output)

			Expect(testUI.Out).To(Say("test-output"))
		})
	})
})

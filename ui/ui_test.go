package ui_test

import (
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
			Expect(testUI.Out).To(Say("This is a test for the UI struct"))
		})
	})

	Describe("DisplayError", func() {
		It("prints text to the err buffer", func() {
			testUI.DisplayError("error text")
			Expect(testUI.Err).To(Say("error text\n"))
		})
	})

	Describe("DisplaySuccess", func() {
		It("prints SUCCESS to the out buffer", func() {
			testUI.DisplaySuccess()
			Expect(testUI.Out).To(Say("SUCCESS"))
		})
	})

	Describe("DisplayFailed", func() {
		It("prints FAILED to the err buffer", func() {
			testUI.DisplayFailed()
			Expect(testUI.Err).To(Say("FAILED"))
		})
	})

	Describe("DisplayNewline", func() {
		It("prints a newline to the out buffer", func() {
			testUI.DisplayNewline()
			Expect(testUI.Out).To(Say("\n"))
		})
	})
})

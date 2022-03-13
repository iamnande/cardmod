package refinements

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter", func() {

	Describe("Source", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "Demi"

				// execution
				actual := NewFilter(expected, "")

				// validations
				Expect(actual.Target()).To(BeEmpty())
				Expect(actual.Source()).To(Equal(expected))

			})
		})
	})

	Describe("Target", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "Quake"

				// execution
				actual := NewFilter("", expected)

				// validations
				Expect(actual.Source()).To(BeEmpty())
				Expect(actual.Target()).To(Equal(expected))

			})
		})
	})

})

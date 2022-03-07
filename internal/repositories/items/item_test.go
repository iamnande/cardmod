package items

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("item", func() {

	Describe("Name", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "M-Stone Piece"

				// setup
				c := &item{name: expected}

				// execution
				actual := c.Name()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

	Describe("Purpose", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "Refinement"

				// setup
				c := &item{purpose: expected}

				// execution
				actual := c.Purpose()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

})

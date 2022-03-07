package magics

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("magic", func() {

	Describe("Name", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "Firaga"

				// setup
				c := &magic{name: expected}

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
				expected := "Offensive"

				// setup
				c := &magic{purpose: expected}

				// execution
				actual := c.Purpose()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

})

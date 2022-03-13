package cards

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("card", func() {

	Describe("Name", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "Abyss Worm"

				// setup
				c := &card{name: expected}

				// execution
				actual := c.Name()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

	Describe("Level", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := int32(5)

				// setup
				c := &card{level: expected}

				// execution
				actual := c.Level()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

})

package limitbreaks

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("limitbreak", func() {

	Describe("Name", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "L?Death"

				// setup
				c := &limitbreak{name: expected}

				// execution
				actual := c.Name()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

})

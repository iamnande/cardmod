package refinements

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("refinement", func() {

	Describe("Source", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "Gesper"

				// setup
				c := &refinement{source: expected}

				// execution
				actual := c.Source()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

	Describe("Target", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := "Degenerator"

				// setup
				c := &refinement{target: expected}

				// execution
				actual := c.Target()

				// validations
				Expect(actual).To(Equal(expected))

			})
		})
	})

	Describe("Numerator/Denominator", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// exepctations
				expected := int32(20)

				// setup
				c := &refinement{
					numerator:   expected,
					denominator: expected,
				}

				// execution
				actualNum, actualDen := c.Numerator(), c.Denominator()

				// validations
				Expect(actualNum).To(Equal(expected))
				Expect(actualDen).To(Equal(expected))

			})
		})
	})

})

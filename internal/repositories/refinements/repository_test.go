package refinements

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {

	Describe("NewRepository", func() {
		Context("When Called", func() {
			It("Should Return", func() {

				// execution
				actual := NewRepository()

				// validations
				Expect(actual).To(BeAssignableToTypeOf(repository{}))
				Expect(len(actual.refinements)).To(Equal(len(refinements)))

				for i := 0; i < len(refinements); i++ {
					Expect(actual.refinements[i].Source()).To(Equal(refinements[i].Source()))
					Expect(actual.refinements[i].Target()).To(Equal(refinements[i].Target()))
					Expect(actual.refinements[i].Numerator()).To(Equal(refinements[i].Numerator()))
					Expect(actual.refinements[i].Denominator()).To(Equal(refinements[i].Denominator()))
				}

			})
		})
	})

	Describe("ListRefinements", func() {
		Context("When Called", func() {
			It("Should Return the Complete Collection", func() {

				// execution
				actual := NewRepository().ListRefinements()

				// validations
				Expect(len(actual)).To(Equal(len(refinements)))
				for i := 0; i < len(refinements); i++ {
					Expect(actual[i].Source()).To(Equal(refinements[i].Source()))
					Expect(actual[i].Target()).To(Equal(refinements[i].Target()))
					Expect(actual[i].Numerator()).To(Equal(refinements[i].Numerator()))
					Expect(actual[i].Denominator()).To(Equal(refinements[i].Denominator()))
				}

			})
		})
	})

	Describe("GetRefinement", func() {
		Context("When an Existing Refinement is Called", func() {
			It("Should Return the Refinement", func() {

				// setup
				i := 5

				// execution
				actual, err := NewRepository().GetRefinement(refinements[i].Source(), refinements[i].Target())

				// validations
				Expect(err).To(BeNil())
				Expect(actual.Source()).To(Equal(refinements[i].Source()))
				Expect(actual.Target()).To(Equal(refinements[i].Target()))

			})
		})
		Context("When a NON-Existing Refinement is Called", func() {
			It("Should NOT Return a Refinement", func() {

				// execution
				actual, err := NewRepository().GetRefinement("no", "no")

				// validations
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("refinement not found"))

			})
		})
	})

})

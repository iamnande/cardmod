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

		// No Filtering
		Context("With No Filters", func() {
			It("Should Return the Complete Collection", func() {

				// execution
				actual := NewRepository().ListRefinements(nil)

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

		// Source Filter ONLY
		// NOTE: This is basically a GetSingular due to the limitation of 1..* relationship of
		// source refinements to target refinements.
		Context("With Source Filter ONLY", func() {
			It("Should Return Source Filtered Results", func() {

				// expectations
				expected := "Tent"

				// execution
				actual := NewRepository().ListRefinements(NewFilter(expected, ""))

				// validations
				Expect(len(actual)).To(Equal(1))
				Expect(actual[0].Source()).To(Equal(expected))

			})
		})

		// Target Filter ONLY
		Context("With Target Filter ONLY", func() {
			It("Should Return Target Filtered Results", func() {

				// expectations
				expected := "Magic Stone"

				// execution
				actual := NewRepository().ListRefinements(NewFilter("", expected))

				// validations
				Expect(len(actual)).To(Equal(3))
				for i := 0; i < 3; i++ {
					Expect(actual[i].Target()).To(Equal(expected))
				}

			})
		})

		// Source AND Target Filter
		// NOTE: This too is essentially a get by source name.
		Context("With Source AND Target Filter", func() {
			It("Should Return Singular Source<->Target Results", func() {

				// expectations
				expectedSource := "Buel"
				expectedTarget := "Magic Stone"

				// execution
				actual := NewRepository().ListRefinements(NewFilter(expectedSource, expectedTarget))

				// validations
				Expect(len(actual)).To(Equal(1))
				Expect(actual[0].Source()).To(Equal(expectedSource))
				Expect(actual[0].Target()).To(Equal(expectedTarget))

			})
		})

	})

})

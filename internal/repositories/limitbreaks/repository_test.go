package limitbreaks

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
				Expect(len(actual.limitbreaks)).To(Equal(len(limitbreaks)))

				for i := 0; i < len(limitbreaks); i++ {
					Expect(actual.limitbreaks[i].Name()).To(Equal(limitbreaks[i].Name()))
				}

			})
		})
	})

	Describe("ListLimitBreaks", func() {
		Context("When Called", func() {
			It("Should Return the Complete Collection", func() {

				// execution
				actual := NewRepository().ListLimitBreaks()

				// validations
				Expect(len(actual)).To(Equal(len(limitbreaks)))
				for i := 0; i < len(limitbreaks); i++ {
					Expect(actual[i].Name()).To(Equal(limitbreaks[i].Name()))
				}

			})
		})
	})

	Describe("GetLimitBreak", func() {
		Context("When an Existing LimitBreak is Called", func() {
			It("Should Return the LimitBreak", func() {

				// setup
				i := 5

				// execution
				actual, err := NewRepository().GetLimitBreak(limitbreaks[i].Name())

				// validations
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(limitbreaks[i].Name()))

			})
		})
		Context("When a NON-Existing LimitBreak is Called", func() {
			It("Should NOT Return a LimitBreak", func() {

				// execution
				actual, err := NewRepository().GetLimitBreak("no")

				// validations
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("limit break not found"))

			})
		})
	})

})

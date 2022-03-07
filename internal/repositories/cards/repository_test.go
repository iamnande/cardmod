package cards

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
				Expect(len(actual.cards)).To(Equal(len(cards)))

				for i := 0; i < len(cards); i++ {
					Expect(actual.cards[i].Name()).To(Equal(cards[i].Name()))
					Expect(actual.cards[i].Level()).To(Equal(cards[i].Level()))
				}

			})
		})
	})

	Describe("ListCards", func() {
		Context("When Called", func() {
			It("Should Return the Complete Collection", func() {

				// execution
				actual := NewRepository().ListCards()

				// validations
				Expect(len(actual)).To(Equal(len(cards)))
				for i := 0; i < len(cards); i++ {
					Expect(actual[i].Name()).To(Equal(cards[i].Name()))
					Expect(actual[i].Level()).To(Equal(cards[i].Level()))
				}

			})
		})
	})

	Describe("GetCard", func() {
		Context("When an Existing Card is Called", func() {
			It("Should Return the Card", func() {

				// setup
				i := 5

				// execution
				actual, err := NewRepository().GetCard(cards[i].Name())

				// validations
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(cards[i].Name()))
				Expect(actual.Level()).To(Equal(cards[i].Level()))

			})
		})
		Context("When a NON-Existing Card is Called", func() {
			It("Should NOT Return a Card", func() {

				// execution
				actual, err := NewRepository().GetCard("no")

				// validations
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("card not found"))

			})
		})
	})

})

package items

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
				Expect(len(actual.items)).To(Equal(len(items)))

				for i := 0; i < len(items); i++ {
					Expect(actual.items[i].Name()).To(Equal(items[i].Name()))
					Expect(actual.items[i].Purpose()).To(Equal(items[i].Purpose()))
				}

			})
		})
	})

	Describe("ListItems", func() {
		Context("When Called", func() {
			It("Should Return the Complete Collection", func() {

				// execution
				actual := NewRepository().ListItems()

				// validations
				Expect(len(actual)).To(Equal(len(items)))
				for i := 0; i < len(items); i++ {
					Expect(actual[i].Name()).To(Equal(items[i].Name()))
					Expect(actual[i].Purpose()).To(Equal(items[i].Purpose()))
				}

			})
		})
	})

	Describe("GetItem", func() {
		Context("When an Existing Item is Called", func() {
			It("Should Return the Item", func() {

				// setup
				i := 5

				// execution
				actual, err := NewRepository().GetItem(items[i].Name())

				// validations
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(items[i].Name()))
				Expect(actual.Purpose()).To(Equal(items[i].Purpose()))

			})
		})
		Context("When a NON-Existing Item is Called", func() {
			It("Should NOT Return a Item", func() {

				// execution
				actual, err := NewRepository().GetItem("no")

				// validations
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("item not found"))

			})
		})
	})

})

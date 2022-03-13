package magics

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
				Expect(len(actual.magics)).To(Equal(len(magics)))

				for i := 0; i < len(magics); i++ {
					Expect(actual.magics[i].Name()).To(Equal(magics[i].Name()))
					Expect(actual.magics[i].Purpose()).To(Equal(magics[i].Purpose()))
				}

			})
		})
	})

	Describe("ListMagics", func() {
		Context("When Called", func() {
			It("Should Return the Complete Collection", func() {

				// execution
				actual := NewRepository().ListMagics()

				// validations
				Expect(len(actual)).To(Equal(len(magics)))
				for i := 0; i < len(magics); i++ {
					Expect(actual[i].Name()).To(Equal(magics[i].Name()))
					Expect(actual[i].Purpose()).To(Equal(magics[i].Purpose()))
				}

			})
		})
	})

	Describe("GetMagic", func() {
		Context("When an Existing Magic is Called", func() {
			It("Should Return the Magic", func() {

				// setup
				i := 5

				// execution
				actual, err := NewRepository().GetMagic(magics[i].Name())

				// validations
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(magics[i].Name()))
				Expect(actual.Purpose()).To(Equal(magics[i].Purpose()))

			})
		})
		Context("When a NON-Existing Magic is Called", func() {
			It("Should NOT Return a Magic", func() {

				// execution
				actual, err := NewRepository().GetMagic("no")

				// validations
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("magic not found"))

			})
		})
	})

})

package errors

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RepositoryError", func() {

	Describe("NewRepositoryError", func() {
		When("Called", func() {
			It("Should Initialize Properly", func() {

				// expectations
				expected := "failure"

				// execution
				err := NewRepositoryError(expected)

				// validation
				Expect(err.Error()).To(Equal(expected))

			})
		})

	})

})

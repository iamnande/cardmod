package cerrors

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
)

var _ = Describe("Errors", func() {

	Describe("APIError", func() {

		// Can properly translate.
		When("Message if 'broken' BaseError is there", func() {
			It("Writes Error String Properly", func() {
				err := &APIError{
					Code:      codes.InvalidArgument,
					Message:   "you whack",
					BaseError: errors.New("failure, Will Robinson"),
				}
				Expect(err.Error()).To(Equal("you whack: failure, Will Robinson"))
			})
		})

	})

	Describe("IsAPIError", func() {

		// Empty condition
		When("It Is nil", func() {
			It("Returns FALSE", func() {
				Expect(IsAPIError(nil)).To(BeFalse())
			})
		})

		// True condition
		When("It Is an APIError", func() {
			It("Returns TRUE", func() {
				err := &APIError{
					Code:      codes.InvalidArgument,
					Message:   "you whack",
					BaseError: errors.New("failure, Will Robinson"),
				}
				Expect(IsAPIError(err)).To(BeTrue())
			})
		})

		// False condition
		When("It Is NOT APIError", func() {
			It("Returns FALSE", func() {
				err := errors.New("nuh uh")
				Expect(IsAPIError(err)).To(BeFalse())
			})
		})

	})

})

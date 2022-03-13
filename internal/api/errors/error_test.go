package errors

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
)

var _ = Describe("APIError", func() {

	Describe("NewAPIError", func() {
		When("Called", func() {
			It("Should Initialize Properly", func() {

				// expectations
				expected := "not found: does not exist"
				expectedCode := codes.NotFound
				expectedMessage := "not found"
				expectedBaseError := errors.New("does not exist")

				// execution
				err := NewAPIError(expectedCode, expectedMessage, expectedBaseError)

				// validation
				Expect(err.Error()).To(Equal(expected))

			})
		})
	})

	Describe("Code", func() {
		When("Called", func() {
			It("Should Return", func() {

				// expectations
				expected := codes.NotFound

				// execution
				err := &APIError{code: expected}

				// validation
				Expect(err.Code()).To(Equal(expected))

			})
		})
	})

	Describe("Message", func() {
		When("Called", func() {
			It("Should Return", func() {

				// expectations
				expected := "not found"

				// execution
				err := &APIError{message: expected}

				// validation
				Expect(err.Message()).To(Equal(expected))

			})
		})
	})

	Describe("BaseError", func() {
		When("Called", func() {
			It("Should Return", func() {

				// expectations
				expected := errors.New("does not exist")

				// execution
				err := &APIError{baseError: expected}

				// validation
				Expect(err.BaseError()).To(Equal(expected))

			})
		})
	})

	Describe("IsAPIError", func() {
		When("It Is", func() {
			It("Should Return TRUE", func() {

				// setup
				err := NewAPIError(codes.Internal, "catastrophic failure", errors.New("database not found"))

				// execution
				actual := IsAPIError(err)

				// validation
				Expect(actual).To(BeTrue())

			})
		})
		When("It Is NOT", func() {
			It("Should Return FALSE", func() {

				// setup
				err := errors.New("does not exist")

				// execution
				actual := IsAPIError(err)

				// validation
				Expect(actual).To(BeFalse())

			})
		})
		When("err is NIL", func() {
			It("Should Return FALSE", func() {

				// execution
				actual := IsAPIError(nil)

				// validation
				Expect(actual).To(BeFalse())

			})
		})
	})

})

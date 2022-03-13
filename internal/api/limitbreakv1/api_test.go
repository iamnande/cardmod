package limitbreakv1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/iamnande/cardmod/pkg/api/limitbreakv1"
)

var _ = Describe("LimitBreakAPI", func() {

	var (
		// ctx for API operations
		ctx context.Context

		// API instance
		api *api
	)

	// Test Case Setup
	BeforeEach(func() {

		// initialize context for API operations
		ctx = context.TODO()

		// initialize API instance
		api = New()

	})

	// GetLimitBreak
	Describe("GetLimitBreak", func() {
		Context("With an Existing LimitBreak", func() {
			It("Should Return the LimitBreak", func() {

				// setup
				req := &limitbreakv1.GetLimitBreakRequest{Name: "Degenerator"}

				// execution
				actual, err := api.GetLimitBreak(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetName()).To(Equal(req.GetName()))

			})
		})
		Context("With a Non-Existing LimitBreak", func() {
			It("Should NOT Return a LimitBreak", func() {

				// setup
				req := &limitbreakv1.GetLimitBreakRequest{Name: "fail"}

				// execution
				actual, err := api.GetLimitBreak(ctx, req)

				// validation
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				// TODO: this feels kind weird but the separate parts make sense
				Expect(err.Error()).To(Equal("limit break not found: limit break not found"))

			})
		})
	})

	// ListLimitBreaks
	Describe("ListLimitBreaks", func() {
		Context("With an Existing LimitBreak", func() {
			It("Should Return the LimitBreak", func() {

				// setup
				req := &limitbreakv1.ListLimitBreaksRequest{}

				// execution
				actual, err := api.ListLimitBreaks(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())

				apiLimitBreaks := actual.GetLimitbreaks()
				dbLimitBreaks := api.limitbreakRepository.ListLimitBreaks()
				for i := 0; i < len(dbLimitBreaks); i++ {
					Expect(apiLimitBreaks[i].GetName()).To(Equal(dbLimitBreaks[i].Name()))
				}

			})
		})
	})

})

package refinementv1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/iamnande/cardmod/pkg/api/refinementv1"
)

var _ = Describe("RefinementAPI", func() {

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

	// GetRefinement
	Describe("GetRefinement", func() {
		Context("With an Existing Refinement", func() {
			It("Should Return the Refinement", func() {

				// setup
				req := &refinementv1.GetRefinementRequest{
					Source: "Tri-Face", Target: "Curse Spike",
				}

				// execution
				actual, err := api.GetRefinement(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetSource()).To(Equal(req.GetSource()))
				Expect(actual.GetTarget()).To(Equal(req.GetTarget()))

			})
		})
		Context("With a Non-Existing Refinement", func() {
			It("Should NOT Return a Refinement", func() {

				// setup
				req := &refinementv1.GetRefinementRequest{
					Source: "fail", Target: "fail",
				}

				// execution
				actual, err := api.GetRefinement(ctx, req)

				// validation
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				// TODO: this feels kind weird but the separate parts make sense
				Expect(err.Error()).To(Equal("refinement not found: refinement not found"))

			})
		})
	})

	// ListRefinements
	Describe("ListRefinements", func() {
		Context("With an Existing Refinement", func() {
			It("Should Return the Refinement", func() {

				// setup
				req := &refinementv1.ListRefinementsRequest{}

				// execution
				actual, err := api.ListRefinements(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())

				apiRefinements := actual.GetRefinements()
				dbRefinements := api.refinementRepository.ListRefinements()
				for i := 0; i < len(dbRefinements); i++ {
					Expect(apiRefinements[i].GetSource()).To(Equal(dbRefinements[i].Source()))
					Expect(apiRefinements[i].GetTarget()).To(Equal(dbRefinements[i].Target()))
					Expect(apiRefinements[i].GetNumerator()).To(Equal(dbRefinements[i].Numerator()))
					Expect(apiRefinements[i].GetDenominator()).To(Equal(dbRefinements[i].Denominator()))
				}

			})
		})
	})

})

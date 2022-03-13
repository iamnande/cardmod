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

	// ListRefinements
	Describe("ListRefinements", func() {
		Context("With no Filtering Supplied", func() {
			It("Should Return the Complete Refinement Collection", func() {

				// setup
				req := &refinementv1.ListRefinementsRequest{}

				// execution
				actual, err := api.ListRefinements(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())

				apiRefinements := actual.GetRefinements()
				dbRefinements := api.refinementRepository.ListRefinements(nil)
				for i := 0; i < len(dbRefinements); i++ {
					Expect(apiRefinements[i].GetSource()).To(Equal(dbRefinements[i].Source()))
					Expect(apiRefinements[i].GetTarget()).To(Equal(dbRefinements[i].Target()))
					Expect(apiRefinements[i].GetNumerator()).To(Equal(dbRefinements[i].Numerator()))
					Expect(apiRefinements[i].GetDenominator()).To(Equal(dbRefinements[i].Denominator()))
				}

			})
		})
		Context("With Target Filtering Supplied", func() {
			It("Should Return the Filtered Target Refinements", func() {

				// setup
				req := &refinementv1.ListRefinementsRequest{
					Filter: &refinementv1.ListRefinementsRequest_Filter{
						Target: "Magic Stone",
					},
				}

				// execution
				actual, err := api.ListRefinements(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetRefinements()).To(HaveLen(3))
				Expect(actual.GetRefinements()[0].GetTarget()).To(Equal(req.GetFilter().GetTarget()))

			})
		})
		Context("With Source AND Target Filtering Supplied", func() {
			It("Should Return the Filtered Refinements", func() {

				// setup
				req := &refinementv1.ListRefinementsRequest{
					Filter: &refinementv1.ListRefinementsRequest_Filter{
						Source: "Buel",
						Target: "Magic Stone",
					},
				}

				// execution
				actual, err := api.ListRefinements(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetRefinements()).To(HaveLen(1))
				Expect(actual.GetRefinements()[0].GetSource()).To(Equal(req.GetFilter().GetSource()))
				Expect(actual.GetRefinements()[0].GetTarget()).To(Equal(req.GetFilter().GetTarget()))

			})
		})
	})

})

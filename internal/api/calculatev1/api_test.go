package calculatev1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/iamnande/cardmod/pkg/api/calculatev1"
)

var _ = Describe("CalculateAPI", func() {

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

	// Calculate
	Describe("Calculate", func() {
		Context("With an Existing Target", func() {
			It("Should Return the Appropriate Refinements", func() {

				// setup
				req := &calculatev1.CalculateRequest{
					Target: "Magic Stone",
					Count:  100,
				}

				// execution
				actual, err := api.Calculate(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetRefinements()).To(HaveLen(3))

			})
		})
		Context("With an NON-Existent Target", func() {
			It("Should Return an Error", func() {

				// setup
				req := &calculatev1.CalculateRequest{
					Target: "Kame House",
					Count:  100,
				}

				// execution
				actual, err := api.Calculate(ctx, req)

				// validation
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With Multiple Child Refinements", func() {
			It("Should Return the Full Refinement Tree", func() {

				// setup
				req := &calculatev1.CalculateRequest{
					Target: "Quake",
					Count:  100,
				}

				// execution
				actual, err := api.Calculate(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetRefinements()).To(HaveLen(1))
				Expect(actual.GetRefinements()[0].Source).To(Equal("Dino Bone"))
				Expect(actual.GetRefinements()[0].Refinements).To(HaveLen(2))
				Expect(actual.GetRefinements()[0].Refinements[0].GetSource()).To(Equal("Armadodo"))
				Expect(actual.GetRefinements()[0].Refinements[1].GetSource()).To(Equal("T-Rexaur"))

			})
		})
	})

})

package cardv1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/iamnande/cardmod/pkg/api/cardv1"
)

var _ = Describe("CardAPI", func() {

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

	// GetCard
	Describe("GetCard", func() {
		Context("With an Existing Card", func() {
			It("Should Return the Card", func() {

				// setup
				req := &cardv1.GetCardRequest{Name: "Quistis"}

				// execution
				actual, err := api.GetCard(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetName()).To(Equal(req.GetName()))

			})
		})
		Context("With a Non-Existing Card", func() {
			It("Should NOT Return a Card", func() {

				// setup
				req := &cardv1.GetCardRequest{Name: "fail"}

				// execution
				actual, err := api.GetCard(ctx, req)

				// validation
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				// TODO: this feels kind weird but the separate parts make sense
				Expect(err.Error()).To(Equal("card not found: card not found"))

			})
		})
	})

	// ListCards
	Describe("ListCards", func() {
		Context("With an Existing Card", func() {
			It("Should Return the Card", func() {

				// setup
				req := &cardv1.ListCardsRequest{}

				// execution
				actual, err := api.ListCards(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())

				apiCards := actual.GetCards()
				dbCards := api.cardRepository.ListCards()
				for i := 0; i < len(dbCards); i++ {
					Expect(apiCards[i].GetName()).To(Equal(dbCards[i].Name()))
					Expect(apiCards[i].GetLevel()).To(Equal(dbCards[i].Level()))
				}

			})
		})
	})

})

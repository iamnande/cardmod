package itemv1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/iamnande/cardmod/pkg/api/itemv1"
)

var _ = Describe("ItemAPI", func() {

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

	// GetItem
	Describe("GetItem", func() {
		Context("With an Existing Item", func() {
			It("Should Return the Item", func() {

				// setup
				req := &itemv1.GetItemRequest{Name: "M-Stone Piece"}

				// execution
				actual, err := api.GetItem(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetName()).To(Equal(req.GetName()))

			})
		})
		Context("With a Non-Existing Item", func() {
			It("Should NOT Return a Item", func() {

				// setup
				req := &itemv1.GetItemRequest{Name: "fail"}

				// execution
				actual, err := api.GetItem(ctx, req)

				// validation
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				// TODO: this feels kind weird but the separate parts make sense
				Expect(err.Error()).To(Equal("item not found: item not found"))

			})
		})
	})

	// ListItems
	Describe("ListItems", func() {
		Context("With an Existing Item", func() {
			It("Should Return the Item", func() {

				// setup
				req := &itemv1.ListItemsRequest{}

				// execution
				actual, err := api.ListItems(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())

				apiItems := actual.GetItems()
				dbItems := api.itemRepository.ListItems()
				for i := 0; i < len(dbItems); i++ {
					Expect(apiItems[i].GetName()).To(Equal(dbItems[i].Name()))
				}

			})
		})
	})

})

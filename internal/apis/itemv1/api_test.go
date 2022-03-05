package itemv1

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"syreclabs.com/go/faker"

	"github.com/iamnande/cardmod/internal/cerrors"
	mockItemRepo "github.com/iamnande/cardmod/internal/daos/mocks"
	"github.com/iamnande/cardmod/internal/models"
	mockItem "github.com/iamnande/cardmod/internal/models/mocks"
	"github.com/iamnande/cardmod/pkg/api/itemv1"
)

var _ = Describe("ItemAPI", func() {

	var (

		// ctx for API and DB operations
		ctx context.Context

		// gomock controller
		ctrl *gomock.Controller

		// mock item repo
		itemDB *mockItemRepo.MockItemDAO

		// itemAPI instance
		itemAPI *api
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize gomock controller
		ctrl = gomock.NewController(GinkgoT())

		// initialize item repository
		itemDB = mockItemRepo.NewMockItemDAO(ctrl)

		// initialize ItemAPI
		itemAPI = New(itemDB)
	})

	// GetItem
	Describe("GetItem", func() {
		Context("With a Proper Name", func() {
			It("Should Get the Item", func() {

				// setup
				expected := "Inferno Fang"

				// init mock item instance
				fakeItem := mockItem.NewMockItem(ctrl)
				fakeItem.EXPECT().Name().Return(expected)

				// mock repository call
				req := &itemv1.GetItemRequest{
					Name: expected,
				}
				itemDB.EXPECT().GetItem(ctx, expected).Return(fakeItem, nil)

				// make the call
				actual, err := itemAPI.GetItem(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetName()).To(Equal(expected))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should NOT Create the Item", func() {

				// mock repository call
				req := &itemv1.GetItemRequest{}
				itemDB.EXPECT().GetItem(ctx, "").Return(nil, &cerrors.APIError{
					Code:      codes.NotFound,
					Message:   "not found",
					BaseError: errors.New("not found"),
				})

				// make the call
				actual, err := itemAPI.GetItem(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With Internal DB Failure", func() {
			It("Should NOT Create the Item", func() {

				// mock repository call
				req := &itemv1.GetItemRequest{}
				itemDB.EXPECT().GetItem(ctx, "").Return(nil, errors.New("not found"))

				// make the call
				actual, err := itemAPI.GetItem(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListItems
	Describe("ListItems", func() {

		// seed the mock repo with items
		var numItems int
		items := make([]models.Item, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numItems = rand.Intn(max-min+1) + min

			for i := 0; i < numItems; i++ {
				item := mockItem.NewMockItem(ctrl)
				item.EXPECT().Name().Return(faker.Name().String()).AnyTimes()
				items = append(items, item)
			}

		})

		Context("When Success", func() {
			It("Should Return a List of Items", func() {

				// mock repository call
				req := &itemv1.ListItemsRequest{}
				itemDB.EXPECT().ListItems(ctx).Return(items, nil)

				// make the call
				actual, err := itemAPI.ListItems(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetItems()[0].GetName()).To(Equal(items[0].Name()))

			})
		})
		Context("When Internal DB Failure Happens", func() {
			It("Should Return an Error", func() {

				// mock repository call
				req := &itemv1.ListItemsRequest{}
				itemDB.EXPECT().ListItems(ctx).Return(nil, errors.New("interal failure"))

				// make the call
				actual, err := itemAPI.ListItems(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

})

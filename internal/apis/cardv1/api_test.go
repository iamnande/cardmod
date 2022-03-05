package cardv1

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
	mockCardRepo "github.com/iamnande/cardmod/internal/daos/mocks"
	"github.com/iamnande/cardmod/internal/models"
	mockCard "github.com/iamnande/cardmod/internal/models/mocks"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
)

var _ = Describe("CardAPI", func() {

	var (

		// ctx for API and DB operations
		ctx context.Context

		// gomock controller
		ctrl *gomock.Controller

		// mock card repo
		cardDB *mockCardRepo.MockCardDAO

		// cardAPI instance
		cardAPI *api
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize gomock controller
		ctrl = gomock.NewController(GinkgoT())

		// initialize card repository
		cardDB = mockCardRepo.NewMockCardDAO(ctrl)

		// initialize CardAPI
		cardAPI = New(cardDB)
	})

	// GetCard
	Describe("GetCard", func() {
		Context("With a Proper Name", func() {
			It("Should Get the Card", func() {

				// setup
				expected := "Tonberry King"

				// init mock card instance
				fakeCard := mockCard.NewMockCard(ctrl)
				fakeCard.EXPECT().Name().Return(expected)

				// mock repository call
				req := &cardv1.GetCardRequest{
					Name: expected,
				}
				cardDB.EXPECT().GetCard(ctx, expected).Return(fakeCard, nil)

				// make the call
				actual, err := cardAPI.GetCard(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetName()).To(Equal(expected))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should NOT Create the Card", func() {

				// mock repository call
				req := &cardv1.GetCardRequest{}
				cardDB.EXPECT().GetCard(ctx, "").Return(nil, &cerrors.APIError{
					Code:      codes.NotFound,
					Message:   "not found",
					BaseError: errors.New("not found"),
				})

				// make the call
				actual, err := cardAPI.GetCard(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With Internal DB Failure", func() {
			It("Should NOT Create the Card", func() {

				// mock repository call
				req := &cardv1.GetCardRequest{}
				cardDB.EXPECT().GetCard(ctx, "").Return(nil, errors.New("not found"))

				// make the call
				actual, err := cardAPI.GetCard(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListCards
	Describe("ListCards", func() {

		// seed the mock repo with cards
		var numCards int
		cards := make([]models.Card, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numCards = rand.Intn(max-min+1) + min

			for i := 0; i < numCards; i++ {
				card := mockCard.NewMockCard(ctrl)
				card.EXPECT().Name().Return(faker.Name().String()).AnyTimes()
				cards = append(cards, card)
			}

		})

		Context("When Success", func() {
			It("Should Return a List of Cards", func() {

				// mock repository call
				req := &cardv1.ListCardsRequest{}
				cardDB.EXPECT().ListCards(ctx).Return(cards, nil)

				// make the call
				actual, err := cardAPI.ListCards(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetCards()[0].GetName()).To(Equal(cards[0].Name()))

			})
		})
		Context("When Internal DB Failure Happens", func() {
			It("Should Return an Error", func() {

				// mock repository call
				req := &cardv1.ListCardsRequest{}
				cardDB.EXPECT().ListCards(ctx).Return(nil, errors.New("interal failure"))

				// make the call
				actual, err := cardAPI.ListCards(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

})

package cardapi

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"syreclabs.com/go/faker"

	"github.com/iamnande/cardmod/internal/domains/card"
	mockCard "github.com/iamnande/cardmod/internal/domains/card/mocks"
	mockCardRepo "github.com/iamnande/cardmod/internal/repositories/mocks"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
)

var _ = Describe("CardAPI", func() {

	var (
		// ctx for operations
		ctx context.Context

		// gomock controller
		ctrl *gomock.Controller

		// mock card repo client
		cardDB *mockCardRepo.MockCardDAO

		// cardAPI instance
		cardAPI api
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize mock controller
		ctrl = gomock.NewController(GinkgoT())

		// initialize card repository instance
		cardDB = mockCardRepo.NewMockCardDAO(ctrl)

		// initialize CardAPI
		cardAPI = New(cardDB)

	})

	// ListCards
	Describe("ListCards", func() {

		// seed the database with cards to be listed
		var numCards int
		cards := make([]card.Card, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numCards = rand.Intn(max-min+1) + min

			for i := 0; i < numCards; i++ {
				card := mockCard.NewMockCard(ctrl)
				card.EXPECT().ID().Return(uuid.New()).AnyTimes()
				card.EXPECT().Name().Return(faker.Name().String()).AnyTimes()
				cards = append(cards, card)
			}

		})

		Context("With no query", func() {
			It("Should return all records", func() {

				// mock repository call
				request := new(cardv1.ListCardsRequest)
				cardDB.EXPECT().
					ListCards(ctx, "").
					Return(cards, nil)

				// retrieve the magic
				actual, err := cardAPI.ListCards(ctx, request)

				// validate expectations
				Expect(err).To(BeNil())
				Expect(actual.GetCards()[0].Id).To(Equal(cards[0].ID().String()))

			})
		})
		Context("With a supplied query", func() {
			It("Should return filtered results", func() {

				// mock repository call
				response := []card.Card{cards[0]}
				request := &cardv1.ListCardsRequest{
					Filter: &cardv1.ListCardsRequest_Filter{
						Query: cards[0].Name(),
					},
				}
				cardDB.EXPECT().
					ListCards(ctx, cards[0].Name()).
					Return(response, nil)

				// retrieve the magic
				actual, err := cardAPI.ListCards(ctx, request)

				// validate expectations
				Expect(err).To(BeNil())
				Expect(actual.GetCards()[0].Id).To(Equal(cards[0].ID().String()))

			})
		})
	})

	// GetCard
	Describe("GetCard", func() {
		Context("With a proper ID, that exists", func() {
			It("Should find the magic", func() {

				// expectations
				name := "Tonberry King"
				uid := uuid.New()

				// init domain mock instance
				fakeCard := mockCard.NewMockCard(ctrl)
				fakeCard.EXPECT().ID().Return(uid)
				fakeCard.EXPECT().Name().Return(name)

				// mock repository call
				request := &cardv1.GetCardRequest{CardId: uid.String()}
				cardDB.EXPECT().
					GetCard(ctx, uid).
					Return(fakeCard, nil)

				// retrieve the magic
				actual, err := cardAPI.GetCard(ctx, request)

				// validate expectations
				Expect(err).To(BeNil())
				Expect(actual.Id).To(Equal(uid.String()))

			})
		})
		Context("With a proper ID, that does not exist", func() {
			It("Should not find the magic", func() {

				// expectations
				uid := uuid.New()

				// mock repository call
				request := &cardv1.GetCardRequest{CardId: uid.String()}
				cardDB.EXPECT().
					GetCard(ctx, uid).
					Return(nil, fmt.Errorf("not found"))

				// retrieve the card
				actual, err := cardAPI.GetCard(ctx, request)

				// validate expectations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With an invalid ID", func() {
			It("Should not fail early, before the lookup", func() {

				// retrieve the card
				actual, err := cardAPI.GetCard(ctx, nil)

				// validate expectations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

})

func TestCardAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CardAPI Suite")
}

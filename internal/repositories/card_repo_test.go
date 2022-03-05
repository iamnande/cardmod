package repositories

import (
	"context"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"syreclabs.com/go/faker"

	"github.com/iamnande/cardmod/internal/database/enttest"
	"github.com/iamnande/cardmod/internal/models"
)

var _ = Describe("Card", func() {

	var (
		// ctx for operation
		ctx context.Context

		// card repo client
		cardDB *cardRepository
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize test database client
		client := enttest.Open(GinkgoT(), "sqlite3", "file:CardRepository?mode=memory&cache=shared&_fk=1")

		// initialize card repository instance
		cardDB = NewCardRepository(client)

	})

	// CreateCard
	Describe("CreateCard", func() {
		Context("With a Proper Name and Level", func() {
			It("Should Create Without Issue", func() {

				// create the card
				name := "Blood Soul"
				level := int32(3)
				actual, err := cardDB.CreateCard(ctx, name, level)

				// expect no error
				Expect(err).To(BeNil())

				// validate construction
				Expect(actual.Name()).To(Equal(name))
				Expect(actual.Level()).To(Equal(level))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should Fail to Create", func() {

				// create the card
				actual, err := cardDB.CreateCard(ctx, "", 1)

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("Without a Proper Level", func() {
			It("Should Fail to Create", func() {

				// create the card
				actual, err := cardDB.CreateCard(ctx, "", -1)

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// GetCard
	Describe("GetCard", func() {
		Context("With an Name that Exists", func() {
			It("Should Find the Card", func() {

				// create a proper card
				name := "Death Claw"
				level := int32(3)
				created, err := cardDB.CreateCard(ctx, name, level)
				Expect(err).To(BeNil())

				// retrieve the card
				actual, err := cardDB.GetCard(ctx, created.Name())
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(created.Name()))

			})
		})
		Context("With a Non-Existent Card", func() {
			It("Should Fail Retrieve Anything", func() {

				// poor attempt to find anything
				actual, err := cardDB.GetCard(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListCards
	Describe("ListCards", func() {

		// seed the database with cards to be listed
		var numCards int
		cards := make([]models.Card, 0)

		BeforeEach(func() {

			min, max := 2, 10
			rand.Seed(time.Now().UnixNano())
			numCards = rand.Intn(max-min+1) + min

			for i := 0; i < numCards; i++ {
				card, err := cardDB.CreateCard(ctx, faker.Name().String(), int32(numCards))
				Expect(err).To(BeNil())
				cards = append(cards, card)
			}

		})

		Context("When Called", func() {
			It("Should Return All Cards", func() {

				// list the cards
				actual, err := cardDB.ListCards(ctx)

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(BeNumerically(">=", numCards))

				found := false
				for i := 0; i < len(actual); i++ {
					for j := 0; j < numCards; j++ {
						if actual[i].Name() == cards[j].Name() {
							found = true
						}
					}
				}
				Expect(found).To(BeTrue())

			})
		})

	})

})

package repositories

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"syreclabs.com/go/faker"

	"github.com/iamnande/cardmod/internal/database/enttest"
	"github.com/iamnande/cardmod/internal/domains/card"
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
				card, err := cardDB.CreateCard(ctx, faker.Name().String())
				Expect(err).To(BeNil())
				cards = append(cards, card)
			}

		})

		Context("With no search provided", func() {
			It("Should return all cards", func() {

				// list the cards
				actual, err := cardDB.ListCards(ctx, "")

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(Equal(numCards))
				Expect(actual[0].ID()).To(Equal(cards[0].ID()))

			})
		})

		Context("With search provided", func() {
			It("Should return matching card(s)", func() {

				// list the cards
				search := cards[0].Name()
				actual, err := cardDB.ListCards(ctx, search)

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(Equal(1))
				Expect(actual[0].Name()).To(Equal(search))

			})
		})

	})

	// CreateCard
	Describe("CreateCard", func() {
		Context("With a proper name", func() {
			It("Should create without issue", func() {

				// create the card
				name := "Blood Soul"
				actual, err := cardDB.CreateCard(ctx, name)

				// expect no error
				Expect(err).To(BeNil())

				// validate construction
				Expect(actual.ID()).NotTo(BeEmpty())
				Expect(actual.Name()).To(Equal(name))

			})
		})
		Context("Without a proper name", func() {
			It("Should fail to create", func() {

				// create the card
				actual, err := cardDB.CreateCard(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// GetCard
	Describe("GetCard", func() {
		Context("With an ID that exists", func() {
			It("Should find the card", func() {

				// create a proper card
				name := "Death Claw"
				created, err := cardDB.CreateCard(ctx, name)
				Expect(err).To(BeNil())

				// retrieve the card
				actual, err := cardDB.GetCard(ctx, created.ID())
				Expect(err).To(BeNil())
				Expect(actual.ID()).To(Equal(created.ID()))

			})
		})
		Context("With a non-existent card", func() {
			It("Should fail retrieve anything", func() {

				// poor attempt to find anything
				actual, err := cardDB.GetCard(ctx, uuid.New())

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// DeleteCard
	Describe("DeleteCard", func() {
		Context("With an ID that exists", func() {
			It("Should delete the card", func() {

				// create a sacrificial card
				name := "Abyss Worm"
				created, err := cardDB.CreateCard(ctx, name)
				Expect(err).To(BeNil())

				// delete the card
				err = cardDB.DeleteCard(ctx, created.ID())
				Expect(err).To(BeNil())

			})
		})
		Context("With a non-existent card", func() {
			It("Should fail to delete", func() {

				// poor attempt to delete
				err := cardDB.DeleteCard(ctx, uuid.New())

				// expect error
				Expect(err).NotTo(BeNil())

			})
		})
	})
})

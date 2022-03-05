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

var _ = Describe("Item", func() {

	var (
		// ctx for operation
		ctx context.Context

		// item repo client
		itemDB *itemRepository
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize test database client
		client := enttest.Open(GinkgoT(), "sqlite3", "file:ItemRepository?mode=memory&cache=shared&_fk=1")

		// initialize item repository instance
		itemDB = NewItemRepository(client)

	})

	// CreateItem
	Describe("CreateItem", func() {
		Context("With a Proper Name", func() {
			It("Should Create Without Issue", func() {

				// create the item
				name := "Inferno Fang"
				actual, err := itemDB.CreateItem(ctx, name)

				// expect no error
				Expect(err).To(BeNil())

				// validate construction
				Expect(actual.Name()).To(Equal(name))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should Fail to Create", func() {

				// create the item
				actual, err := itemDB.CreateItem(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// GetItem
	Describe("GetItem", func() {
		Context("With an Name That Exists", func() {
			It("Should Find the Item", func() {

				// create a proper item
				name := "Death Claw"
				created, err := itemDB.CreateItem(ctx, name)
				Expect(err).To(BeNil())

				// retrieve the item
				actual, err := itemDB.GetItem(ctx, created.Name())
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(created.Name()))

			})
		})
		Context("With a Non-Existent Item", func() {
			It("Should Fail Retrieve Anything", func() {

				// poor attempt to find anything
				actual, err := itemDB.GetItem(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListItems
	Describe("ListItems", func() {

		// seed the database with items to be listed
		var numItems int
		items := make([]models.Item, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numItems = rand.Intn(max-min+1) + min

			for i := 0; i < numItems; i++ {
				item, err := itemDB.CreateItem(ctx, faker.Name().String())
				Expect(err).To(BeNil())
				items = append(items, item)
			}

		})

		Context("When Called", func() {
			It("Should Return All Items", func() {

				// list the items
				actual, err := itemDB.ListItems(ctx)

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(BeNumerically(">=", numItems))

				found := false
				for i := 0; i < len(actual); i++ {
					for j := 0; j < numItems; j++ {
						if actual[i].Name() == items[j].Name() {
							found = true
						}
					}
				}
				Expect(found).To(BeTrue())

			})
		})

	})

})

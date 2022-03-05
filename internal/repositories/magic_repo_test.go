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

var _ = Describe("Magic", func() {

	var (
		// ctx for operation
		ctx context.Context

		// magic repo client
		magicDB *magicRepository
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize test database client
		client := enttest.Open(GinkgoT(), "sqlite3", "file:MagicRepository?mode=memory&cache=shared&_fk=1")

		// initialize magic repository instance
		magicDB = NewMagicRepository(client)

	})

	// CreateMagic
	Describe("CreateMagic", func() {
		Context("With a Proper Name", func() {
			It("Should Create Without Issue", func() {

				// create the magic
				name := "Firaga"
				purpose := "Offensive"
				actual, err := magicDB.CreateMagic(ctx, name, purpose)

				// expect no error
				Expect(err).To(BeNil())

				// validate construction
				Expect(actual.Name()).To(Equal(name))
				Expect(actual.Purpose()).To(Equal(purpose))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should Fail to Create", func() {

				// create the magic
				actual, err := magicDB.CreateMagic(ctx, "", "Offensive")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("Without a Proper Purpose", func() {
			It("Should Fail to Create", func() {

				// create the magic
				actual, err := magicDB.CreateMagic(ctx, "Firaga", "potatoes")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// GetMagic
	Describe("GetMagic", func() {
		Context("With an Name That Exists", func() {
			It("Should Find the Magic", func() {

				// create a proper magic
				name := "Blizzard"
				created, err := magicDB.CreateMagic(ctx, name, "Indirect")
				Expect(err).To(BeNil())

				// retrieve the magic
				actual, err := magicDB.GetMagic(ctx, created.Name())
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(created.Name()))

			})
		})
		Context("With a Non-Existent Magic", func() {
			It("Should Fail Retrieve Anything", func() {

				// poor attempt to find anything
				actual, err := magicDB.GetMagic(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListMagics
	Describe("ListMagics", func() {

		// seed the database with magics to be listed
		var numMagics int
		magics := make([]models.Magic, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numMagics = rand.Intn(max-min+1) + min

			for i := 0; i < numMagics; i++ {
				magic, err := magicDB.CreateMagic(ctx, faker.Name().String(), "Restorative")
				Expect(err).To(BeNil())
				magics = append(magics, magic)
			}

		})

		Context("When Called", func() {
			It("Should Return All Magics", func() {

				// list the magics
				actual, err := magicDB.ListMagics(ctx)

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(BeNumerically(">=", numMagics))

				found := false
				for i := 0; i < len(actual); i++ {
					for j := 0; j < numMagics; j++ {
						if actual[i].Name() == magics[j].Name() {
							found = true
						}
					}
				}
				Expect(found).To(BeTrue())

			})
		})

	})

})

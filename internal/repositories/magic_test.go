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
	"github.com/iamnande/cardmod/internal/domains/magic"
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

	// ListMagics
	Describe("ListMagics", func() {

		// seed the database with magics to be listed
		var numMagics int
		magics := make([]magic.Magic, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numMagics = rand.Intn(max-min+1) + min

			for i := 0; i < numMagics; i++ {
				magic, err := magicDB.CreateMagic(ctx, faker.Name().String())
				Expect(err).To(BeNil())
				magics = append(magics, magic)
			}

		})

		Context("With no search provided", func() {
			It("Should return all magics", func() {

				// list the magics
				actual, err := magicDB.ListMagics(ctx, "")

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(Equal(numMagics))
				Expect(actual[0].ID()).To(Equal(magics[0].ID()))

			})
		})

		Context("With search provided", func() {
			It("Should return matching magic(s)", func() {

				// list the magics
				search := magics[0].Name()
				actual, err := magicDB.ListMagics(ctx, search)

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(Equal(1))
				Expect(actual[0].Name()).To(Equal(search))

			})
		})

	})

	// CreateMagic
	Describe("CreateMagic", func() {
		Context("With a proper name", func() {
			It("Should create without issue", func() {

				// create the magic
				name := "Flare"
				actual, err := magicDB.CreateMagic(ctx, name)

				// expect no error
				Expect(err).To(BeNil())

				// validate construction
				Expect(actual.ID()).NotTo(BeEmpty())
				Expect(actual.Name()).To(Equal(name))

			})
		})
		Context("Without a proper name", func() {
			It("Should fail to create", func() {

				// create the magic
				actual, err := magicDB.CreateMagic(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// GetMagic
	Describe("GetMagic", func() {
		Context("With an ID that exists", func() {
			It("Should find the magic", func() {

				// create a proper magic
				name := "Tornado"
				created, err := magicDB.CreateMagic(ctx, name)
				Expect(err).To(BeNil())

				// retrieve the magic
				actual, err := magicDB.GetMagic(ctx, created.ID())
				Expect(err).To(BeNil())
				Expect(actual.ID()).To(Equal(created.ID()))

			})
		})
		Context("With a non-existent magic", func() {
			It("Should fail retrieve anything", func() {

				// poor attempt to find anything
				actual, err := magicDB.GetMagic(ctx, uuid.New())

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// DeleteMagic
	Describe("DeleteMagic", func() {
		Context("With an ID that exists", func() {
			It("Should delete the magic", func() {

				// create a sacrificial magic
				name := "Quake"
				created, err := magicDB.CreateMagic(ctx, name)
				Expect(err).To(BeNil())

				// delete the magic
				err = magicDB.DeleteMagic(ctx, created.ID())
				Expect(err).To(BeNil())

			})
		})
		Context("With a non-existent magic", func() {
			It("Should fail to delete", func() {

				// poor attempt to delete
				err := magicDB.DeleteMagic(ctx, uuid.New())

				// expect error
				Expect(err).NotTo(BeNil())

			})
		})
	})
})

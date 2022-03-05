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

var _ = Describe("LimitBreak", func() {

	var (
		// ctx for operation
		ctx context.Context

		// limitbreak repo client
		limitbreakDB *limitbreakRepository
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize test database client
		client := enttest.Open(GinkgoT(), "sqlite3", "file:LimitBreakRepository?mode=memory&cache=shared&_fk=1")

		// initialize limitbreak repository instance
		limitbreakDB = NewLimitBreakRepository(client)

	})

	// CreateLimitBreak
	Describe("CreateLimitBreak", func() {
		Context("With a Proper Name", func() {
			It("Should Create Without Issue", func() {

				// create the limitbreak
				name := "Inferno Fang"
				actual, err := limitbreakDB.CreateLimitBreak(ctx, name)

				// expect no error
				Expect(err).To(BeNil())

				// validate construction
				Expect(actual.Name()).To(Equal(name))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should Fail to Create", func() {

				// create the limitbreak
				actual, err := limitbreakDB.CreateLimitBreak(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// GetLimitBreak
	Describe("GetLimitBreak", func() {
		Context("With an Name That Exists", func() {
			It("Should Find the LimitBreak", func() {

				// create a proper limitbreak
				name := "L?Death"
				created, err := limitbreakDB.CreateLimitBreak(ctx, name)
				Expect(err).To(BeNil())

				// retrieve the limitbreak
				actual, err := limitbreakDB.GetLimitBreak(ctx, created.Name())
				Expect(err).To(BeNil())
				Expect(actual.Name()).To(Equal(created.Name()))

			})
		})
		Context("With a Non-Existent LimitBreak", func() {
			It("Should Fail Retrieve Anything", func() {

				// poor attempt to find anything
				actual, err := limitbreakDB.GetLimitBreak(ctx, "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListLimitBreaks
	Describe("ListLimitBreaks", func() {

		// seed the database with limitbreaks to be listed
		var numLimitBreaks int
		limitbreaks := make([]models.LimitBreak, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numLimitBreaks = rand.Intn(max-min+1) + min

			for i := 0; i < numLimitBreaks; i++ {
				limitbreak, err := limitbreakDB.CreateLimitBreak(ctx, faker.Name().String())
				Expect(err).To(BeNil())
				limitbreaks = append(limitbreaks, limitbreak)
			}

		})

		Context("When Called", func() {
			It("Should Return All LimitBreaks", func() {

				// list the limitbreaks
				actual, err := limitbreakDB.ListLimitBreaks(ctx)

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(BeNumerically(">=", numLimitBreaks))

				found := false
				for i := 0; i < len(actual); i++ {
					for j := 0; j < numLimitBreaks; j++ {
						if actual[i].Name() == limitbreaks[j].Name() {
							found = true
						}
					}
				}
				Expect(found).To(BeTrue())

			})
		})

	})

})

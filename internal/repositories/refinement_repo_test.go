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

var _ = Describe("Refinement", func() {

	var (
		// ctx for operation
		ctx context.Context

		// refinement repo client
		refinementDB *refinementRepository
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize test database client
		client := enttest.Open(GinkgoT(), "sqlite3", "file:RefinementRepository?mode=memory&cache=shared&_fk=1")

		// initialize refinement repository instance
		refinementDB = NewRefinementRepository(client)

	})

	// CreateRefinement
	Describe("CreateRefinement", func() {
		Context("With a Proper Source and Target", func() {
			It("Should Create Without Issue", func() {

				// create the refinement
				source := "Abyss Worm"
				target := "Tornado"
				numerator := int32(1)
				denominator := int32(20)
				actual, err := refinementDB.CreateRefinement(ctx, source, target, numerator, denominator)

				// expect no error
				Expect(err).To(BeNil())

				// validate construction
				Expect(actual.Source()).To(Equal(source))
				Expect(actual.Target()).To(Equal(target))
				Expect(actual.Numerator()).To(Equal(numerator))
				Expect(actual.Denominator()).To(Equal(denominator))

			})
		})
		Context("Without a Proper Source", func() {
			It("Should Fail to Create", func() {

				// create the refinement
				target := "Tornado"
				numerator := int32(1)
				denominator := int32(20)
				actual, err := refinementDB.CreateRefinement(ctx, "", target, numerator, denominator)

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("Without a Proper Target", func() {
			It("Should Fail to Create", func() {

				// create the refinement
				source := "Abyss Worm"
				numerator := int32(1)
				denominator := int32(20)
				actual, err := refinementDB.CreateRefinement(ctx, source, "", numerator, denominator)

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// GetRefinement
	Describe("GetRefinement", func() {
		Context("With an Source and Target That Exists", func() {
			It("Should Find the Refinement", func() {

				// create a proper refinement
				source := "Gesper"
				target := "Black Hole"
				numerator := int32(1)
				denominator := int32(20)
				created, err := refinementDB.CreateRefinement(ctx, source, target, numerator, denominator)
				Expect(err).To(BeNil())

				// retrieve the refinement
				actual, err := refinementDB.GetRefinement(ctx, created.Source(), created.Target())
				Expect(err).To(BeNil())
				Expect(actual.Source()).To(Equal(created.Source()))
				Expect(actual.Target()).To(Equal(created.Target()))

			})
		})
		Context("With a Non-Existent Refinement", func() {
			It("Should Fail Retrieve Anything", func() {

				// poor attempt to find anything
				actual, err := refinementDB.GetRefinement(ctx, "", "")

				// expect error
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListRefinements
	Describe("ListRefinements", func() {

		// seed the database with refinements to be listed
		var numRefinements int
		refinements := make([]models.Refinement, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numRefinements = rand.Intn(max-min+1) + min

			for i := 0; i < numRefinements; i++ {
				refinement, err := refinementDB.CreateRefinement(ctx,
					faker.Name().String(), faker.Name().String(), 1, 20)
				Expect(err).To(BeNil())
				refinements = append(refinements, refinement)
			}

		})

		Context("When Called", func() {
			It("Should Return All Refinements", func() {

				// list the refinements
				actual, err := refinementDB.ListRefinements(ctx)

				// expect no error
				Expect(err).To(BeNil())

				// validate expectations
				Expect(len(actual)).To(BeNumerically(">=", numRefinements))

				found := false
				for i := 0; i < len(actual); i++ {
					for j := 0; j < numRefinements; j++ {
						if actual[i].Source() == refinements[j].Source() &&
							actual[i].Target() == refinements[j].Target() {
							found = true
						}
					}
				}
				Expect(found).To(BeTrue())

			})
		})

	})

})

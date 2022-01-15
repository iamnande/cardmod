package magicapi

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

	"github.com/iamnande/cardmod/internal/domains/magic"
	mockMagic "github.com/iamnande/cardmod/internal/domains/magic/mocks"
	mockMagicRepo "github.com/iamnande/cardmod/internal/repositories/mocks"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
)

var _ = Describe("MagicAPI", func() {

	var (
		// ctx for operations
		ctx context.Context

		// gomock controller
		ctrl *gomock.Controller

		// mock magic repo client
		magicDB *mockMagicRepo.MockMagicDAO

		// magicAPI instance
		magicAPI api
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize mock controller
		ctrl = gomock.NewController(GinkgoT())

		// initialize magic repository instance
		magicDB = mockMagicRepo.NewMockMagicDAO(ctrl)

		// initialize MagicAPI
		magicAPI = New(magicDB)

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
				magic := mockMagic.NewMockMagic(ctrl)
				magic.EXPECT().ID().Return(uuid.New()).AnyTimes()
				magic.EXPECT().Name().Return(faker.Name().String()).AnyTimes()
				magics = append(magics, magic)
			}

		})

		Context("With no query", func() {
			It("Should return all records", func() {

				// mock repository call
				request := new(magicv1.ListMagicsRequest)
				magicDB.EXPECT().
					ListMagics(ctx, "").
					Return(magics, nil)

				// retrieve the magic
				actual, err := magicAPI.ListMagics(ctx, request)

				// validate expectations
				Expect(err).To(BeNil())
				Expect(actual.GetMagics()[0].Id).To(Equal(magics[0].ID().String()))

			})
		})
		Context("With a supplied query", func() {
			It("Should return filtered results", func() {

				// mock repository call
				response := []magic.Magic{magics[0]}
				request := &magicv1.ListMagicsRequest{
					Filter: &magicv1.ListMagicsRequest_Filter{
						Query: magics[0].Name(),
					},
				}
				magicDB.EXPECT().
					ListMagics(ctx, magics[0].Name()).
					Return(response, nil)

				// retrieve the magic
				actual, err := magicAPI.ListMagics(ctx, request)

				// validate expectations
				Expect(err).To(BeNil())
				Expect(actual.GetMagics()[0].Id).To(Equal(magics[0].ID().String()))

			})
		})
	})

	// GetMagic
	Describe("GetMagic", func() {
		Context("With a proper ID, that exists", func() {
			It("Should find the magic", func() {

				// expectations
				name := "Tonberry King"
				uid := uuid.New()

				// init domain mock instance
				fakeMagic := mockMagic.NewMockMagic(ctrl)
				fakeMagic.EXPECT().ID().Return(uid)
				fakeMagic.EXPECT().Name().Return(name)

				// mock repository call
				request := &magicv1.GetMagicRequest{MagicId: uid.String()}
				magicDB.EXPECT().
					GetMagic(ctx, uid).
					Return(fakeMagic, nil)

				// retrieve the magic
				actual, err := magicAPI.GetMagic(ctx, request)

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
				request := &magicv1.GetMagicRequest{MagicId: uid.String()}
				magicDB.EXPECT().
					GetMagic(ctx, uid).
					Return(nil, fmt.Errorf("not found"))

				// retrieve the magic
				actual, err := magicAPI.GetMagic(ctx, request)

				// validate expectations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With an invalid ID", func() {
			It("Should not fail early, before the lookup", func() {

				// retrieve the magic
				actual, err := magicAPI.GetMagic(ctx, nil)

				// validate expectations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

})

func TestMagicAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MagicAPI Suite")
}

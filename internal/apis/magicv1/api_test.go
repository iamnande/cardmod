package magicv1

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
	mockMagicRepo "github.com/iamnande/cardmod/internal/daos/mocks"
	"github.com/iamnande/cardmod/internal/models"
	mockMagic "github.com/iamnande/cardmod/internal/models/mocks"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
)

var _ = Describe("MagicAPI", func() {

	var (

		// ctx for API and DB operations
		ctx context.Context

		// gomock controller
		ctrl *gomock.Controller

		// mock magic repo
		magicDB *mockMagicRepo.MockMagicDAO

		// magicAPI instance
		magicAPI *api
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize gomock controller
		ctrl = gomock.NewController(GinkgoT())

		// initialize magic repository
		magicDB = mockMagicRepo.NewMockMagicDAO(ctrl)

		// initialize MagicAPI
		magicAPI = New(magicDB)
	})

	// GetMagic
	Describe("GetMagic", func() {
		Context("With a Proper Name", func() {
			It("Should Get the Magic", func() {

				// setup
				expectedName := "Firaga"
				expectedPurpose := "Offensive"

				// init mock magic instance
				fakeMagic := mockMagic.NewMockMagic(ctrl)
				fakeMagic.EXPECT().Name().Return(expectedName)
				fakeMagic.EXPECT().Purpose().Return(expectedPurpose)

				// mock repository call
				req := &magicv1.GetMagicRequest{
					Name: expectedName,
				}
				magicDB.EXPECT().GetMagic(ctx, expectedName).Return(fakeMagic, nil)

				// make the call
				actual, err := magicAPI.GetMagic(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetName()).To(Equal(expectedName))
				Expect(actual.GetPurpose()).To(Equal(expectedPurpose))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should NOT Create the Magic", func() {

				// mock repository call
				req := &magicv1.GetMagicRequest{}
				magicDB.EXPECT().GetMagic(ctx, "").Return(nil, &cerrors.APIError{
					Code:      codes.NotFound,
					Message:   "not found",
					BaseError: errors.New("not found"),
				})

				// make the call
				actual, err := magicAPI.GetMagic(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With Internal DB Failure", func() {
			It("Should NOT Create the Magic", func() {

				// mock repository call
				req := &magicv1.GetMagicRequest{}
				magicDB.EXPECT().GetMagic(ctx, "").Return(nil, errors.New("not found"))

				// make the call
				actual, err := magicAPI.GetMagic(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListMagics
	Describe("ListMagics", func() {

		// seed the mock repo with magics
		var numMagics int
		magics := make([]models.Magic, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numMagics = rand.Intn(max-min+1) + min

			for i := 0; i < numMagics; i++ {
				magic := mockMagic.NewMockMagic(ctrl)
				magic.EXPECT().Name().Return(faker.Name().String()).AnyTimes()
				magic.EXPECT().Purpose().Return("Indirect").AnyTimes()
				magics = append(magics, magic)
			}

		})

		Context("When Success", func() {
			It("Should Return a List of Magics", func() {

				// mock repository call
				req := &magicv1.ListMagicsRequest{}
				magicDB.EXPECT().ListMagics(ctx).Return(magics, nil)

				// make the call
				actual, err := magicAPI.ListMagics(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetMagics()[0].GetName()).To(Equal(magics[0].Name()))

			})
		})
		Context("When Internal DB Failure Happens", func() {
			It("Should Return an Error", func() {

				// mock repository call
				req := &magicv1.ListMagicsRequest{}
				magicDB.EXPECT().ListMagics(ctx).Return(nil, errors.New("interal failure"))

				// make the call
				actual, err := magicAPI.ListMagics(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

})

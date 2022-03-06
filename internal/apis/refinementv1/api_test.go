package refinementv1

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
	mockRefinementRepo "github.com/iamnande/cardmod/internal/daos/mocks"
	"github.com/iamnande/cardmod/internal/models"
	mockRefinement "github.com/iamnande/cardmod/internal/models/mocks"
	"github.com/iamnande/cardmod/pkg/api/refinementv1"
)

var _ = Describe("RefinementAPI", func() {

	var (

		// ctx for API and DB operations
		ctx context.Context

		// gomock controller
		ctrl *gomock.Controller

		// mock refinement repo
		refinementDB *mockRefinementRepo.MockRefinementDAO

		// refinementAPI instance
		refinementAPI *api
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize gomock controller
		ctrl = gomock.NewController(GinkgoT())

		// initialize refinement repository
		refinementDB = mockRefinementRepo.NewMockRefinementDAO(ctrl)

		// initialize RefinementAPI
		refinementAPI = New(refinementDB)
	})

	// GetRefinement
	Describe("GetRefinement", func() {
		Context("With a Proper Source and Target", func() {
			It("Should Get the Refinement", func() {

				// setup
				source := "Ruby Dragon"
				target := "Inferno Fang"
				numerator := int32(10)
				denominator := int32(1)

				// init mock refinement instance
				fakeRefinement := mockRefinement.NewMockRefinement(ctrl)
				fakeRefinement.EXPECT().Source().Return(source)
				fakeRefinement.EXPECT().Target().Return(target)
				fakeRefinement.EXPECT().Numerator().Return(numerator)
				fakeRefinement.EXPECT().Denominator().Return(denominator)

				// mock repository call
				req := &refinementv1.GetRefinementRequest{
					Source: source,
					Target: target,
				}
				refinementDB.EXPECT().GetRefinement(ctx, source, target).Return(fakeRefinement, nil)

				// make the call
				actual, err := refinementAPI.GetRefinement(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetSource()).To(Equal(source))
				Expect(actual.GetTarget()).To(Equal(target))

			})
		})
		Context("Without a Proper Source or Target", func() {
			It("Should NOT Create the Refinement", func() {

				// mock repository call
				req := &refinementv1.GetRefinementRequest{}
				refinementDB.EXPECT().GetRefinement(ctx, "", "").Return(nil, &cerrors.APIError{
					Code:      codes.NotFound,
					Message:   "not found",
					BaseError: errors.New("not found"),
				})

				// make the call
				actual, err := refinementAPI.GetRefinement(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With Internal DB Failure", func() {
			It("Should NOT Create the Refinement", func() {

				// mock repository call
				req := &refinementv1.GetRefinementRequest{}
				refinementDB.EXPECT().GetRefinement(ctx, "", "").Return(nil, errors.New("not found"))

				// make the call
				actual, err := refinementAPI.GetRefinement(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListRefinements
	Describe("ListRefinements", func() {

		// seed the mock repo with refinements
		var numRefinements int
		refinements := make([]models.Refinement, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numRefinements = rand.Intn(max-min+1) + min

			for i := 0; i < numRefinements; i++ {
				refinement := mockRefinement.NewMockRefinement(ctrl)
				refinement.EXPECT().Source().Return(faker.Name().String()).AnyTimes()
				refinement.EXPECT().Target().Return(faker.Name().String()).AnyTimes()
				refinement.EXPECT().Numerator().Return(int32(10)).AnyTimes()
				refinement.EXPECT().Denominator().Return(int32(10)).AnyTimes()
				refinements = append(refinements, refinement)
			}

		})

		Context("When Success", func() {
			It("Should Return a List of Refinements", func() {

				// mock repository call
				req := &refinementv1.ListRefinementsRequest{}
				refinementDB.EXPECT().ListRefinements(ctx).Return(refinements, nil)

				// make the call
				actual, err := refinementAPI.ListRefinements(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetRefinements()[0].GetSource()).To(Equal(refinements[0].Source()))

			})
		})
		Context("When Internal DB Failure Happens", func() {
			It("Should Return an Error", func() {

				// mock repository call
				req := &refinementv1.ListRefinementsRequest{}
				refinementDB.EXPECT().ListRefinements(ctx).Return(nil, errors.New("interal failure"))

				// make the call
				actual, err := refinementAPI.ListRefinements(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

})

package limitbreakv1

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
	mockLimitBreakRepo "github.com/iamnande/cardmod/internal/daos/mocks"
	"github.com/iamnande/cardmod/internal/models"
	mockLimitBreak "github.com/iamnande/cardmod/internal/models/mocks"
	"github.com/iamnande/cardmod/pkg/api/limitbreakv1"
)

var _ = Describe("LimitBreakAPI", func() {

	var (

		// ctx for API and DB operations
		ctx context.Context

		// gomock controller
		ctrl *gomock.Controller

		// mock limitbreak repo
		limitbreakDB *mockLimitBreakRepo.MockLimitBreakDAO

		// limitbreakAPI instance
		limitbreakAPI *api
	)

	BeforeEach(func() {

		// initialize context
		ctx = context.Background()

		// initialize gomock controller
		ctrl = gomock.NewController(GinkgoT())

		// initialize limitbreak repository
		limitbreakDB = mockLimitBreakRepo.NewMockLimitBreakDAO(ctrl)

		// initialize LimitBreakAPI
		limitbreakAPI = New(limitbreakDB)
	})

	// GetLimitBreak
	Describe("GetLimitBreak", func() {
		Context("With a Proper Name", func() {
			It("Should Get the LimitBreak", func() {

				// setup
				expected := "L?Death"

				// init mock limitbreak instance
				fakeLimitBreak := mockLimitBreak.NewMockLimitBreak(ctrl)
				fakeLimitBreak.EXPECT().Name().Return(expected)

				// mock repository call
				req := &limitbreakv1.GetLimitBreakRequest{
					Name: expected,
				}
				limitbreakDB.EXPECT().GetLimitBreak(ctx, expected).Return(fakeLimitBreak, nil)

				// make the call
				actual, err := limitbreakAPI.GetLimitBreak(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetName()).To(Equal(expected))

			})
		})
		Context("Without a Proper Name", func() {
			It("Should NOT Create the LimitBreak", func() {

				// mock repository call
				req := &limitbreakv1.GetLimitBreakRequest{}
				limitbreakDB.EXPECT().GetLimitBreak(ctx, "").Return(nil, &cerrors.APIError{
					Code:      codes.NotFound,
					Message:   "not found",
					BaseError: errors.New("not found"),
				})

				// make the call
				actual, err := limitbreakAPI.GetLimitBreak(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
		Context("With Internal DB Failure", func() {
			It("Should NOT Create the LimitBreak", func() {

				// mock repository call
				req := &limitbreakv1.GetLimitBreakRequest{}
				limitbreakDB.EXPECT().GetLimitBreak(ctx, "").Return(nil, errors.New("not found"))

				// make the call
				actual, err := limitbreakAPI.GetLimitBreak(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

	// ListLimitBreaks
	Describe("ListLimitBreaks", func() {

		// seed the mock repo with limitbreaks
		var numLimitBreaks int
		limitbreaks := make([]models.LimitBreak, 0)

		BeforeEach(func() {

			min, max := 3, 15
			rand.Seed(time.Now().UnixNano())
			numLimitBreaks = rand.Intn(max-min+1) + min

			for i := 0; i < numLimitBreaks; i++ {
				limitbreak := mockLimitBreak.NewMockLimitBreak(ctrl)
				limitbreak.EXPECT().Name().Return(faker.Name().String()).AnyTimes()
				limitbreaks = append(limitbreaks, limitbreak)
			}

		})

		Context("When Success", func() {
			It("Should Return a List of LimitBreaks", func() {

				// mock repository call
				req := &limitbreakv1.ListLimitBreaksRequest{}
				limitbreakDB.EXPECT().ListLimitBreaks(ctx).Return(limitbreaks, nil)

				// make the call
				actual, err := limitbreakAPI.ListLimitBreaks(ctx, req)

				// validate expecations
				Expect(err).To(BeNil())
				Expect(actual.GetLimitbreaks()[0].GetName()).To(Equal(limitbreaks[0].Name()))

			})
		})
		Context("When Internal DB Failure Happens", func() {
			It("Should Return an Error", func() {

				// mock repository call
				req := &limitbreakv1.ListLimitBreaksRequest{}
				limitbreakDB.EXPECT().ListLimitBreaks(ctx).Return(nil, errors.New("interal failure"))

				// make the call
				actual, err := limitbreakAPI.ListLimitBreaks(ctx, req)

				// validate expecations
				Expect(err).NotTo(BeNil())
				Expect(actual).To(BeNil())

			})
		})
	})

})

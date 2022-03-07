package magicv1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/iamnande/cardmod/pkg/api/magicv1"
)

var _ = Describe("MagicAPI", func() {

	var (
		// ctx for API operations
		ctx context.Context

		// API instance
		api *api
	)

	// Test Case Setup
	BeforeEach(func() {

		// initialize context for API operations
		ctx = context.TODO()

		// initialize API instance
		api = New()

	})

	// GetMagic
	Describe("GetMagic", func() {
		Context("With an Existing Magic", func() {
			It("Should Return the Magic", func() {

				// setup
				req := &magicv1.GetMagicRequest{Name: "Blizzara"}

				// execution
				actual, err := api.GetMagic(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())
				Expect(actual.GetName()).To(Equal(req.GetName()))

			})
		})
		Context("With a Non-Existing Magic", func() {
			It("Should NOT Return a Magic", func() {

				// setup
				req := &magicv1.GetMagicRequest{Name: "fail"}

				// execution
				actual, err := api.GetMagic(ctx, req)

				// validation
				Expect(actual).To(BeNil())
				Expect(err).NotTo(BeNil())
				// TODO: this feels kind weird but the separate parts make sense
				Expect(err.Error()).To(Equal("magic not found: magic not found"))

			})
		})
	})

	// ListMagics
	Describe("ListMagics", func() {
		Context("With an Existing Magic", func() {
			It("Should Return the Magic", func() {

				// setup
				req := &magicv1.ListMagicsRequest{}

				// execution
				actual, err := api.ListMagics(ctx, req)

				// validation
				Expect(err).To(BeNil())
				Expect(actual).NotTo(BeNil())

				apiMagics := actual.GetMagics()
				dbMagics := api.magicRepository.ListMagics()
				for i := 0; i < len(dbMagics); i++ {
					Expect(apiMagics[i].GetName()).To(Equal(dbMagics[i].Name()))
					Expect(apiMagics[i].GetPurpose()).To(Equal(dbMagics[i].Purpose()))
				}

			})
		})
	})

})

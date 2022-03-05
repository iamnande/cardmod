package config

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Describe("mustLoad", func() {

		// Default Test Cases
		Context("Default Values", func() {

			// No env vars supplied
			When("No Environment Variables are Supplied", func() {
				It("Loads All Defaults", func() {
					actual := mustLoad()
					Expect(actual.Environment).To(Equal(defaultEnvironment))
					Expect(actual.Server.Port).To(Equal(defaultServerPort))
					Expect(actual.Server.ShutdownGracePeriod).To(Equal(defaultServerShutdownGracePeriod))
				})
			})

			// Some env vars supplied
			When("Some Environment Variables are Supplied", func() {
				It("Injects Supplied & Defaults the Rest", func() {
					expected := "test"
					_ = os.Setenv("CARDMOD_ENVIRONMENT", expected)
					actual := mustLoad()
					Expect(actual.Environment).To(Equal(expected))
					Expect(actual.Server.Port).To(Equal(defaultServerPort))
					Expect(actual.Server.ShutdownGracePeriod).To(Equal(defaultServerShutdownGracePeriod))
				})
			})

		})

	})

	Describe("MustLoad", func() {

		// Is only loaded truly once
		When("Called Multiple Times", func() {
			It("Uses Once-ed/Cached Config", func() {
				expected := MustLoad()
				actual := MustLoad()
				Expect(actual).To(Equal(expected))
			})
		})

	})

})

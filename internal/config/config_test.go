package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Describe("PortListener", func() {

		// Can properly translate.
		When("Port is 9000", func() {
			It("Translates to ':9000'", func() {
				cfg := &ServerConfig{
					Port: 9000,
				}
				actual := cfg.PortListener()
				Expect(actual).To(Equal(":9000"))
			})
		})

	})

})

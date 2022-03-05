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

	Describe("DSN", func() {

		// Can properly translate.
		When("Values are Provide", func() {
			It("Translates Properly", func() {
				cfg := &DatabaseConfig{
					Hostname: "hostname",
					Port:     12345,
					Name:     "name",
					SSLMode:  "prefer",
					Username: "username",
					Password: "password",
				}
				actual := cfg.DSN()
				Expect(actual).To(Equal("postgres://hostname:12345/name?sslmode=prefer&user=username&password=password"))
			})
		})

	})

})

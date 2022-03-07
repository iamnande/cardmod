package errors

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAPIErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Errors Suite")
}

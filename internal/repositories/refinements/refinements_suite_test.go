package refinements

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRefinements(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Refinements Suite")
}

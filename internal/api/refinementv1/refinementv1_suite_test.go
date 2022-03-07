package refinementv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRefinementv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Refinementv1 Suite")
}

package limitbreakv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLimitBreakv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LimitBreakv1 Suite")
}

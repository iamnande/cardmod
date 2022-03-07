package limitbreaks

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLimitBreaks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LimitBreaks Suite")
}

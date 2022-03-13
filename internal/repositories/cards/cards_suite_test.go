package cards

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCards(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cards Suite")
}

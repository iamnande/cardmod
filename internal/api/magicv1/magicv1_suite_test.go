package magicv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMagicv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Magicv1 Suite")
}

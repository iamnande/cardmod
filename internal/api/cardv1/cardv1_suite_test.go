package cardv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCardv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cardv1 Suite")
}

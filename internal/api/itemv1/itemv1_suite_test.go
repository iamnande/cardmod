package itemv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestItemv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Itemv1 Suite")
}

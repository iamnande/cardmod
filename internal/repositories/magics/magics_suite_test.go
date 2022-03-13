package magics

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMagics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Magics Suite")
}

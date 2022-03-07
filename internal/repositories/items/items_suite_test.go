package items

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestItems(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Items Suite")
}

package calculatev1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCalculatev1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculatev1 Suite")
}

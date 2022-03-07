package errors

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepositoryErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Errors Suite")
}

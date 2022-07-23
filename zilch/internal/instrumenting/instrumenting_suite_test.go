package instrumenting_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInstrumenting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Instrumenting Suite")
}

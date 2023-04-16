package instrumenting_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInstrumenting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Instrumenting Suite")
}

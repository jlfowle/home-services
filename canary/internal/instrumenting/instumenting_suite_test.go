package instrumenting_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInstumenting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Instumenting Suite")
}

package instrumenting_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInstumenting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Instumenting Suite")
}

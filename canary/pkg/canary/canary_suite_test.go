package canary_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCanary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Canary Suite")
}

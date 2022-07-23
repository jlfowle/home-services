package zilch_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestZilch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zilch Suite")
}

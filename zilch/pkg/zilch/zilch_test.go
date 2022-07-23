package zilch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

var _ = Describe("Zilch", func() {
	When("it is used", func() {
		It("succeeds", func() {
			svc := zilch.NewZilchService()
			Expect(svc.Get()).To(BeNil())
		})
	})
})

package canary_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/canary/pkg/canary"
)

var (
	svc canary.CanaryService
)

var _ = Describe("Canary Service", func() {
	BeforeEach(func() {
		svc = canary.NewCanaryService()
	})
	It("should be healthy", func() {
		resp, err := svc.GetServiceHealth()
		Expect(err).ToNot(HaveOccurred())
		Expect(resp).To(BeEquivalentTo(200))
	})
	It("should be ready", func() {
		resp, err := svc.GetServiceReadiness()
		Expect(err).ToNot(HaveOccurred())
		Expect(resp).To(BeEquivalentTo(200))
	})
})

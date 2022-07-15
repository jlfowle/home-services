package endpoints_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/canary/pkg/canary"
	"github.com/jlfowle/home-services/canary/pkg/canary/endpoints"
)

var (
	eps endpoints.Set
	ctx context.Context
)

var _ = Describe("Endpoints", func() {
	var _ = BeforeEach(func() {
		svc := canary.NewService()
		eps = endpoints.NewEndpointSet(svc)
		ctx = context.TODO()
	})
	Describe("Get", func() {
		It("Is successful", func() {
			Expect(eps.Get(ctx)).Should(Succeed())
		})
	})
	Describe("ServiceStatus", func() {
		It("Returns a success", func() {
			res, err := eps.ServiceStatus(ctx)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).To(BeEquivalentTo(200))
		})
	})
})

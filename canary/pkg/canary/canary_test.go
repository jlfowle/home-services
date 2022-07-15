package canary_test

import (
	"context"

	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/canary/pkg/canary"
)

var ctx context.Context
var _ = Describe("Canary", func() {
	var _ = BeforeEach(func() {
		ctx = context.TODO()
	})
	Describe("Get", func() {
		It("Is successful", func() {
			err := canary.NewService().Get(ctx)
			Expect(err).ToNot(HaveOccurred())
		})
	})
	Describe("Service Status", func() {
		It("Is successful", func() {
			status, err := canary.NewService().ServiceStatus(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(status).To(Equal(http.StatusOK))
		})
	})
})

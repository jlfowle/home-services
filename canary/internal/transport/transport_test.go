package transport_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/jlfowle/home-services/canary/internal/transport"
)

type dummyService struct{}

func (s dummyService) GetServiceHealth() (int, error) {
	return 200, nil
}

func (s dummyService) GetServiceReadiness() (int, error) {
	return 200, nil
}

var _ = Describe("Transport", func() {
	var server *ghttp.Server
	BeforeEach(func() {
		svc := &dummyService{}
		server = ghttp.NewServer()
		handler := transport.MakeGetEndpointHandler(svc)
		server.AppendHandlers(
			http.HandlerFunc(handler.ServeHTTP),
		)
	})
	It("responds to get health requests", func() {
		resp, err := http.Get(server.URL() + "/healthz")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.StatusCode).To(BeEquivalentTo(200))
	})
})

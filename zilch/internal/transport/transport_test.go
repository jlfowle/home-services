package transport_test

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/jlfowle/home-services/zilch/internal/transport"
)

type dummyService struct{}

func (s dummyService) Get() error {
	return nil
}

type mockSet struct{}

func (s *mockSet) GetEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return nil, nil
	}
}

var _ = Describe("Transport", func() {
	var server *ghttp.Server
	BeforeEach(func() {
		// start a test http server
		handler := transport.MakeGetEndpointHandler(&mockSet{})
		server = ghttp.NewServer()
		server.AppendHandlers(
			http.HandlerFunc(handler.ServeHTTP),
		)
	})
	AfterEach(func() {
		server.Close()
	})
	When("It is created", func() {
		It("It serves http", func() {
			resp, err := http.Get(server.URL() + "/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
	})
})

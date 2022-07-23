package transport_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/jlfowle/home-services/zilch/internal/transport"
)

type dummyService struct{}

func (s dummyService) Get() error {
	return nil
}

var _ = Describe("Transport", func() {
	var server *ghttp.Server
	BeforeEach(func() {
		// start a test http server
		svc := &dummyService{}
		server = ghttp.NewServer()
		handler := transport.MakeGetEndpointHandler(svc)
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

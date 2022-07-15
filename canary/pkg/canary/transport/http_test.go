package transport_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/jlfowle/home-services/canary/pkg/canary"
	"github.com/jlfowle/home-services/canary/pkg/canary/endpoints"
	"github.com/jlfowle/home-services/canary/pkg/canary/transport"
)

var _ = Describe("Server", func() {
	var server *ghttp.Server
	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})
	AfterEach(func() {
		server.Close()
	})
	Context("When get request is sent to empty path", func() {
		BeforeEach(func() {
			svc := canary.NewService()
			eps := endpoints.NewEndpointSet(svc)
			handler := transport.NewHTTPHandler(eps)
			server.AppendHandlers(
				http.HandlerFunc(handler.ServeHTTP),
			)
		})
		It("Is responds ok", func() {
			resp, err := http.Get(server.URL() + "/get")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
		It("Is healthy", func() {
			resp, err := http.Get(server.URL() + "/healthz")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).To(Equal("{\"status\":200}\n"))
		})
	})
})

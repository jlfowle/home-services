package endpoints_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/zilch/internal/endpoints"
	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

type mockService struct{}

var testError error

func (s *mockService) Get() error {
	return testError
}

var _ = Describe("endpoints", func() {
	var eps endpoints.Set
	var svc zilch.ZilchService
	BeforeEach(func() {
		svc = &mockService{}
		eps = endpoints.NewEndpointSet(svc)
		testError = fmt.Errorf("this is the error")
	})
	It("returns and endpoint", func() {
		ep := eps.GetEndpoint()
		_, err := ep(nil, nil)
		Expect(err).To(BeIdenticalTo(testError))
	})
})

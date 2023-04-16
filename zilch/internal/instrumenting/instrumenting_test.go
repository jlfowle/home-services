package instrumenting_test

import (
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/zilch/internal/instrumenting"
	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

type dummyService struct{}

func (s dummyService) Get() error {
	return nil
}

var (
	svc             zilch.ZilchService
	countCV         *stdprometheus.CounterVec
	countMetric     metrics.Counter
	histogramHV     *stdprometheus.SummaryVec
	histogramMetric metrics.Histogram
	fieldKeys       []string
)

var _ = Describe("Instrumenting", func() {
	BeforeEach(func() {
		fieldKeys = []string{}
		countCV = stdprometheus.NewCounterVec(stdprometheus.CounterOpts{
			Namespace: "http",
			Subsystem: "request",
			Name:      "total",
			Help:      "Number of requests recieved.",
		}, fieldKeys)
		countMetric = kitprometheus.NewCounter(countCV)
		histogramHV = stdprometheus.NewSummaryVec(stdprometheus.SummaryOpts{
			Namespace: "http",
			Subsystem: "request",
			Name:      "duration_seconds",
			Help:      "Total duration of requests in seconds.",
		}, fieldKeys)
		histogramMetric = kitprometheus.NewSummary(histogramHV)
	})
	When("A function is called", func() {
		It("Increments the metrics", func() {
			svc = &dummyService{}
			svc = instrumenting.NewInstrumentingMiddleware(countMetric, histogramMetric, svc)
			Expect(svc.Get()).To(Succeed())
			Expect(testutil.CollectAndCount(countCV)).To(BeEquivalentTo(1))
			Expect(testutil.CollectAndCount(histogramHV)).To(BeEquivalentTo(1))
		})
	})
})

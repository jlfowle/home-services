package instrumenting

import (
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/jlfowle/home-services/canary/pkg/canary"
)

func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, svc canary.CanaryService) canary.CanaryService {
	return &instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		next:           svc,
	}
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           canary.CanaryService
}

func (mw instrumentingMiddleware) GetServiceHealth() (n int, err error) {
	defer func(begin time.Time) {
		lvs := []string{}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(float64(time.Since(begin).Seconds()))
	}(time.Now())

	n, err = mw.next.GetServiceHealth()
	return
}

func (mw instrumentingMiddleware) GetServiceReadiness() (n int, err error) {
	defer func(begin time.Time) {
		lvs := []string{}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(float64(time.Since(begin).Seconds()))
	}(time.Now())

	n, err = mw.next.GetServiceReadiness()
	return
}

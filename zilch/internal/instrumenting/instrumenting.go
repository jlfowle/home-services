package instrumenting

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, svc zilch.ZilchService) zilch.ZilchService {
	return &instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		next:           svc,
	}
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           zilch.ZilchService
}

func (mw instrumentingMiddleware) Get() (err error) {
	defer func(begin time.Time) {
		lvs := []string{}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(float64(time.Since(begin).Seconds()))
	}(time.Now())

	err = mw.next.Get()
	return
}

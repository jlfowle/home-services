package logging

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/jlfowle/home-services/canary/pkg/canary"
)

func NewLoggingMiddleware(logger log.Logger, svc canary.CanaryService) canary.CanaryService {
	return &loggingMiddleware{logger: logger, next: svc}
}

type loggingMiddleware struct {
	logger log.Logger
	next   canary.CanaryService
}

func (mw loggingMiddleware) GetServiceHealth() (n int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get",
			"err", err,
			"status", n,
			"duration", time.Since(begin),
		)
	}(time.Now())

	n, err = mw.next.GetServiceHealth()
	return
}

func (mw loggingMiddleware) GetServiceReadiness() (n int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get",
			"err", err,
			"status", n,
			"duration", time.Since(begin),
		)
	}(time.Now())

	n, err = mw.next.GetServiceReadiness()
	return
}

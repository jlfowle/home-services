package logging

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

func NewLoggingMiddleware(logger log.Logger, svc zilch.ZilchService) zilch.ZilchService {
	return &loggingMiddleware{logger: logger, next: svc}
}

type loggingMiddleware struct {
	logger log.Logger
	next   zilch.ZilchService
}

func (mw loggingMiddleware) Get() (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get",
			"err", err,
			"duration", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.Get()
	return
}

package logging_test

import (
	"github.com/go-kit/kit/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/zilch/internal/logging"
	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

type dummyWriter struct {
	Logs [][]byte
}

func (d *dummyWriter) Write(p []byte) (n int, err error) {
	d.Logs = append(d.Logs, p)
	return len(p), nil
}

type dummyService struct{}

func (s dummyService) Get() error {
	return nil
}

var (
	svc    zilch.ZilchService
	logger log.Logger
)

var _ = Describe("Logging", func() {
	When("The middleware is called", func() {
		It("Logs a message", func() {
			dummyLog := &dummyWriter{}
			logger = log.NewLogfmtLogger(dummyLog)
			svc = &dummyService{}
			svc = logging.NewLoggingMiddleware(logger, svc)
			svc.Get()
			Expect(len(dummyLog.Logs)).To(BeEquivalentTo(1))
		})
	})
})

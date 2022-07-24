package logging_test

import (
	"github.com/go-kit/kit/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/canary/internal/logging"
	"github.com/jlfowle/home-services/canary/pkg/canary"
)

type dummyWriter struct {
	Logs [][]byte
}

func (d *dummyWriter) Write(p []byte) (n int, err error) {
	d.Logs = append(d.Logs, p)
	return len(p), nil
}

type dummyService struct{}

func (s dummyService) GetServiceHealth() (int, error) {
	return 200, nil
}

func (s dummyService) GetServiceReadiness() (int, error) {
	return 200, nil
}

var (
	svc      canary.CanaryService
	dummyLog *dummyWriter
	logger   log.Logger
)

var _ = Describe("Logging", func() {
	BeforeEach(func() {
		dummyLog = &dummyWriter{}
		logger = log.NewLogfmtLogger(dummyLog)
		svc = &dummyService{}
		svc = logging.NewLoggingMiddleware(logger, svc)
	})
	Describe("Get Service Heath", func() {
		It("logs a message", func() {
			_, _ = svc.GetServiceHealth()
			Expect(len(dummyLog.Logs)).To(BeEquivalentTo(1))
		})
	})
	Describe("Get Service Readiness", func() {
		It("logs a message", func() {
			_, _ = svc.GetServiceReadiness()
			Expect(len(dummyLog.Logs)).To(BeEquivalentTo(1))
		})
	})
})

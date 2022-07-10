package canary

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

type canaryService struct{}

func NewService() Service { return &canaryService{} }

func (c *canaryService) Get(_ context.Context) error {
	return nil
}

func (c *canaryService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("msg", "Checking the Service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

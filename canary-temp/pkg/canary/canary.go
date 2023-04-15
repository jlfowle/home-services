package canary

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/log"
)

type canaryService struct{}

func NewService() Service { return &canaryService{} }

func (c *canaryService) Ping(_ context.Context) string {
	return "Pong"
}

func (w *canaryService) Health(_ context.Context) (int, error) {
	_ = logger.Log("msg", "Checking the health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

package canary

import (
	"context"
)

type Service interface {
	// Respond to request with Pong
	Ping(ctx context.Context) string
	Health(ctx context.Context) (int, error)
}

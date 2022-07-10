package canary

import (
	"context"
)

type Service interface {
	// Get the list of all documents
	Get(ctx context.Context) error
	ServiceStatus(ctx context.Context) (int, error)
}

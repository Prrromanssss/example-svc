package clickhouse

import (
	"context"

	"example-svc/internal/some_domain/entity"

	"github.com/google/uuid"
)

type SomeClient struct {
}

func NewSomeClient() *SomeClient {
	return &SomeClient{}
}

func (s *SomeClient) CreateEvent(ctx context.Context, event entity.Event) (uuid.UUID, error) {
	return uuid.New(), nil
}

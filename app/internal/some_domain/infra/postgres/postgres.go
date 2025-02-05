package postgres

import (
	"context"
	"example-svc/internal/some_domain/entity"
	"example-svc/internal/some_domain/usecases/models"

	"github.com/google/uuid"
)

type SomeClient struct {
}

func NewSomeClient() *SomeClient {
	return &SomeClient{}
}

func (s *SomeClient) CreateSomeEntity(ctx context.Context, entity entity.SomeEntity) (int64, error) {
	return 0, nil
}

func (s *SomeClient) AcquireEntityByEventID(ctx context.Context, eventID uuid.UUID) (entity.SomeEntity, error) {
	return entity.NewSomeEntity("1", "2", "3"), nil
}

func (s *SomeClient) GetSomeModel(ctx context.Context, model models.GetExampleQuery) (models.ExampleModel, error) {
	return models.ExampleModel{}, nil
}

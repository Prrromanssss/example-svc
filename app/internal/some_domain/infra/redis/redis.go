package redis

import (
	"context"
	"example-svc/internal/some_domain/usecases/models"
)

type SomeClient struct {
}

func NewSomeClient() *SomeClient {
	return &SomeClient{}
}

func (s *SomeClient) GetSomeValue(ctx context.Context, key string) (models.ExampleModel, error) {
	return models.ExampleModel{}, nil
}

func (s *SomeClient) SetSomeValue(ctx context.Context, key string, value models.ExampleModel) error {
	return nil
}

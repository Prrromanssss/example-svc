package delivery

import (
	"context"
	"example-svc/internal/some_domain/usecases/models"

	"github.com/google/uuid"
)

type SomeDomain interface {
	SomeMethodCreate(ctx context.Context, command models.CreateCommand) (int64, error)
	GetChatIDByUserID(ctx context.Context, userID int64) (int64, error)
	ProcessSomeEvent(ctx context.Context, eventID uuid.UUID) (uuid.UUID, error)
	GetExample(ctx context.Context, query models.GetExampleQuery) (models.ExampleModel, error)
	CloseExpiredSession(ctx context.Context) error
}

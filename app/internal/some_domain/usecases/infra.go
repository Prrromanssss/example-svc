// Secondary adapters.
package usecases

import (
	"context"
	"example-svc/internal/some_domain/entity"
	"example-svc/internal/some_domain/usecases/models"

	"github.com/google/uuid"
)

type AuthorizationManager interface {
	CanAccess(methodID int) (bool, error)
}

type SomeStorage interface {
	CreateSomeEntity(ctx context.Context, entity entity.SomeEntity) (int64, error)
	AcquireEntityByEventID(ctx context.Context, eventID uuid.UUID) (entity.SomeEntity, error)
	GetSomeModel(ctx context.Context, model models.GetExampleQuery) (models.ExampleModel, error)
}

type UserManager interface {
	GetChatIDByUserID(ctx context.Context, userID int64) (int64, error)
}

type SomeEventStorage interface {
	CreateEvent(ctx context.Context, event entity.Event) (uuid.UUID, error)
}

type SomeCache interface {
	GetSomeValue(ctx context.Context, key string) (models.ExampleModel, error)
	SetSomeValue(ctx context.Context, key string, value models.ExampleModel) error
}

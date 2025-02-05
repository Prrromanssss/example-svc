package usecases

import (
	"context"
	"errors"
	"example-svc/internal/some_domain/entity"
	"example-svc/internal/some_domain/usecases/models"

	"github.com/google/uuid"
)

type SomeExampleUsecase struct {
	authorizationManager AuthorizationManager
	someEntityStorage    SomeStorage
	someEventStorage     SomeEventStorage
	cacheStorage         SomeCache
	userManager          UserManager
}

func NewSomeExampleUsecase(
	authorizationManager AuthorizationManager,
	someEntityStorage SomeStorage,
	someEventStorage SomeEventStorage,
	cacheStorage SomeCache,
	userManager UserManager,
) *SomeExampleUsecase {
	return &SomeExampleUsecase{
		authorizationManager: authorizationManager,
		someEntityStorage:    someEntityStorage,
		someEventStorage:     someEventStorage,
		cacheStorage:         cacheStorage,
		userManager:          userManager,
	}
}

func (u *SomeExampleUsecase) SomeMethodCreate(ctx context.Context, command models.CreateCommand) (int64, error) {
	access, err := u.authorizationManager.CanAccess(SomeMethodID)
	if err != nil {
		return 0, err
	}

	if !access {
		return 0, ErrNoAccess
	}

	entity1, err := entity.CreateSomeEntity(
		command.Param1,
		command.Param2,
		command.Param3,
	)
	if err != nil {
		return 0, err
	}

	entityID, err := u.someEntityStorage.CreateSomeEntity(ctx, entity1)
	if err != nil {
		return 0, err
	}

	return entityID, nil
}

func (u *SomeExampleUsecase) GetChatIDByUserID(ctx context.Context, userID int64) (int64, error) {
	return u.userManager.GetChatIDByUserID(ctx, userID)
}

func (u *SomeExampleUsecase) ProcessSomeEvent(ctx context.Context, eventID uuid.UUID) (uuid.UUID, error) {
	entity1, err := u.someEntityStorage.AcquireEntityByEventID(ctx, eventID)
	if err != nil {
		return [16]byte{}, err
	}

	processedEvent, err := entity1.ProcessSomeEvent()
	if err != nil {
		return [16]byte{}, err
	}

	processedID, err := u.someEventStorage.CreateEvent(ctx, processedEvent)
	if err != nil {
		return [16]byte{}, err
	}

	return processedID, nil
}

func (u *SomeExampleUsecase) GetExample(ctx context.Context, query models.GetExampleQuery) (models.ExampleModel, error) {
	value, err := u.cacheStorage.GetSomeValue(ctx, query.Param1)
	if err != nil {
		if errors.Is(err, ErrCacheNotFound) {
			model, err := u.someEntityStorage.GetSomeModel(ctx, query)
			if err != nil {
				return models.ExampleModel{}, err
			}

			err = u.cacheStorage.SetSomeValue(ctx, query.Param1, model)
			if err != nil {
				return models.ExampleModel{}, err
			}

			return model, nil
		}
	}

	return value, nil
}

func (u *SomeExampleUsecase) CloseExpiredSession(ctx context.Context) error {
	return nil
}

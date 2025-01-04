package usecase

import (
	"context"
	"pitch_backend/domain/model"
	"pitch_backend/domain/repository"

	"github.com/samber/do"
)

type GameUsecase interface {
	Create(ctx context.Context, game *model.Game) error
}

type GameUsecaseImpl struct {
	gameRepository repository.GameRepository
}

func NewGameUsecase(i *do.Injector) (GameUsecase, error) {
	gameRepository := do.MustInvoke[repository.GameRepository](i)

	return &GameUsecaseImpl{
		gameRepository: gameRepository,
	}, nil
}

func (u GameUsecaseImpl) Create(ctx context.Context, game *model.Game) error {
	err := u.gameRepository.Create(ctx, game)
	if err != nil {
		return err
	}
	return nil
}

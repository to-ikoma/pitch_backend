package repository

import (
	"context"
	"pitch_backend/domain/model"
)

type GameRepository interface {
	Create(ctx context.Context, game *model.Game) error
}

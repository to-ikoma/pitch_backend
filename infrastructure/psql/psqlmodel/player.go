package psqlmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Player struct {
	bun.BaseModel     `bun:"table:player,alias:player"`
	ID                uuid.UUID `bun:"id,pk"`
	Name              string    `bun:"name"`
	FavoritePositions []string  `bun:"position"`
	TeamID            uuid.UUID `bun:"team_id"`
	Team              *Team     `bun:"rel:belongs-to, join:team_id=id"`
	UserID            uuid.UUID `bun:"user_id"`
	CreatedAt         time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt         time.Time `bun:"updated_at,default:current_timestamp"`
}

package psqlmodel

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Player struct {
	bun.BaseModel     `bun:"table:player,alias:player"`
	ID                uuid.UUID       `bun:"id,pk"`
	Name              string          `bun:"name"`
	FavoritePositions []string        `bun:"position"`
	TeamID            uuid.UUID       `bun:"team_id"`
	Team              *Team           `bun:"rel:belongs-to, join:team_id=id"`
	StartingMember    *StartingMember `bun:"rel:has-many, join:id=player_id"`
}

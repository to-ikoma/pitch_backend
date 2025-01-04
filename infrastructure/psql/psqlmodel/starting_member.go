package psqlmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type StartingMember struct {
	bun.BaseModel `bun:"table:starting_member,alias:starting_member"`
	ID            uuid.UUID `bun:"id,pk"`
	LineupID      uuid.UUID `bun:"lineup_id"`
	Lineup        *Lineup   `bun:"rel:belongs-to, join:lineup_id=id"`
	PlayerID      uuid.UUID `bun:"player_id"`
	Player        *Player   `bun:"rel:belongs-to, join:player_id=id"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}

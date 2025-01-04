package psqlmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Lineup struct {
	bun.BaseModel   `bun:"table:lineup,alias:lineup"`
	ID              uuid.UUID         `bun:"id,pk"`
	GameID          uuid.UUID         `bun:"game_id"`
	Game            *Game             `bun:"rel:belongs-to,join:game_id=id"`
	StartingMembers []*StartingMember `bun:"rel:has-many,join:id=lineup_id"`
	TeamID          uuid.UUID         `bun:"team_id"`
	Team            *Team             `bun:"rel:belongs-to,join:team_id=id"`
	CreatedAt       time.Time         `bun:"created_at,default:current_timestamp"`
	UpdatedAt       time.Time         `bun:"updated_at,default:current_timestamp"`
}

package psqlmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Team struct {
	bun.BaseModel `bun:"table:team,alias:team"`
	ID            uuid.UUID `bun:"id,pk"`
	Name          string    `bun:"name"`
	Players       []*Player `bun:"rel:has-many,join:id=team_id"`
	Lineups       []*Lineup `bun:"rel:has-many,join:id=team_id"`
	UserID        uuid.UUID `bun:"user_id"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}

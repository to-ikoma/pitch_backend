package psqlmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Game struct {
	bun.BaseModel `bun:"table:game,alias:game"`
	ID            uuid.UUID `bun:"id,pk"`
	Date          time.Time `bun:"date"`
	Title         string    `bun:"title"`
	IsEnableDH    bool      `bun:"is_enable_dh"`
	Lineups       []*Lineup `bun:"rel:has-many,join:id=game_id"`
	UserID        uuid.UUID `bun:"user_id"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}

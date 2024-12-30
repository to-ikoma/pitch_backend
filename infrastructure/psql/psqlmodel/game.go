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
}

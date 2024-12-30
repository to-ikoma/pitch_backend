package psqlmodel

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Team struct {
	bun.BaseModel `bun:"table:team,alias:team"`
	ID            uuid.UUID `bun:"id,pk"`
	Name          string    `bun:"name"`
	Players       []*Player `bun:"rel:has-many,join:id=team_id"`
	Lineups       []*Lineup `bun:"rel:has-many,join:id=team_id"`
}

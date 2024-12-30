package model

import "github.com/google/uuid"

type LineupID uuid.UUID

func (i LineupID) String() string {
	return uuid.UUID(i).String()
}

func NewLineupID() LineupID {
	return LineupID(uuid.New())
}

func NewLineupIDFromString(s string) LineupID {
	return LineupID(uuid.MustParse(s))
}

func NewLineupIDFromUUID(u uuid.UUID) LineupID {
	return LineupID(u)
}

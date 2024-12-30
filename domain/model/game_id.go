package model

import "github.com/google/uuid"

type GameID uuid.UUID

func (i GameID) String() string {
	return uuid.UUID(i).String()
}

func NewGameID() GameID {
	return GameID(uuid.New())
}

func NewGameIDFromString(s string) GameID {
	return GameID(uuid.MustParse(s))
}

func NewGameIDFromUUID(id uuid.UUID) GameID {
	return GameID(id)
}

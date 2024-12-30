package model

import "github.com/google/uuid"

type PlayerID uuid.UUID

func (i PlayerID) String() string {
	return uuid.UUID(i).String()
}

func (i PlayerID) UUID() uuid.UUID {
	return uuid.UUID(i)
}

func NewPlayerID() PlayerID {
	return PlayerID(uuid.New())
}

func NewPlayerIDFromString(s string) PlayerID {
	return PlayerID(uuid.MustParse(s))
}

func NewPlayerIDFromUUID(id uuid.UUID) PlayerID {
	return PlayerID(id)
}

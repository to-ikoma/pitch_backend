package model

import "github.com/google/uuid"

type TeamID uuid.UUID

func (i TeamID) String() string {
	return uuid.UUID(i).String()
}

func NewTeamID() TeamID {
	return TeamID(uuid.New())
}

func NewTeamIDFromString(s string) TeamID {
	return TeamID(uuid.MustParse(s))
}

func NewTeamIDFromUUID(id uuid.UUID) TeamID {
	return TeamID(id)
}

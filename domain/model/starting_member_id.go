package model

import "github.com/google/uuid"

type StartingMemberID uuid.UUID

func (s StartingMemberID) String() string {
	return uuid.UUID(s).String()
}

func NewStartingMemberID() StartingMemberID {
	return StartingMemberID(uuid.New())
}

func NewStaringMemberIDFromString(s string) StartingMemberID {
	return StartingMemberID(uuid.MustParse(s))
}

func NewStaringMemberIDFromUUID(u uuid.UUID) StartingMemberID {
	return StartingMemberID(u)
}

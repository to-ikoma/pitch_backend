package model

type StartingMember struct {
	ID          StartingMemberID
	OrderNumber int
	Player      Player
	Position    Position
}

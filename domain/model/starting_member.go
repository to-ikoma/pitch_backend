package model

type StartingMember struct {
	ID          StartingMemberID
	OrderNumber int
	Player      *Player
	Position    Position
}

// type OptinalStatingMember struct {
// 	Player *Player
// }

// func NewStartingMembers(optinal OptinalStatingMember) *StartingMember {
// 	return &StartingMember{
// 		ID:          NewStartingMemberID(),
// 		OrderNumber: 1,
// 		Player:      optinal.Player,
// 	}

// }

// // non DH

// // enable DH

package model

type Player struct {
	ID                PlayerID
	Name              string
	FavoritePositions []Position
}

// func NewPlayer() *Player {
// 	return &Player{
// 		ID:   NewPlayerID(),
// 		Name: "",
// 		FavoritePositions: []Position{
// 			Unspecified,
// 		},
// 	}
// }

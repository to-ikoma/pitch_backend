package model

type Team struct {
	ID      TeamID
	Name    string
	Players []*Player
}

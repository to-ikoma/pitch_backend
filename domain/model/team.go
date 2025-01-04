package model

type Team struct {
	ID      TeamID
	Name    string
	Players []*Player
}

// type OptionalTeam struct {
// 	Players []*Player
// }

// func NewTeam(option OptionalTeam) *Team {
// 	return &Team{
// 		ID:      NewTeamID(),
// 		Name:    "",
// 		Players: option.Players,
// 	}
// }

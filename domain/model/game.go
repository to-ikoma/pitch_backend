package model

type Game struct {
	ID               GameID
	Date             *Date
	Title            string
	IsEnableDH       bool
	FirstTeamLineup  *Lineup
	SecondTeamLineup *Lineup
}

// type OptionalNewGame struct {
// 	Date             *Date
// 	Title            string
// 	IsEnableDH       bool
// 	FirstTeamLineup  *Lineup
// 	SecondTeamLineup *Lineup
// }

// func NewGame(option *OptionalNewGame) *Game {
// 	var newDate *Date
// 	if option.Date != nil {
// 		newDate = NewDate()
// 	}

// 	return &Game{
// 		ID:               NewGameID(),
// 		Date:             newDate,
// 		Title:            option.Title,
// 		IsEnableDH:       option.IsEnableDH,
// 		FirstTeamLineup:  option.FirstTeamLineup,
// 		SecondTeamLineup: option.SecondTeamLineup,
// 	}
// }

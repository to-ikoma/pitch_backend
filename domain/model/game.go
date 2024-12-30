package model

type Game struct {
	ID               GameID
	Date             Date
	Title            string
	IsEnableDH       bool
	FirstTeamLineup  Lineup
	SecondTeamLineup Lineup
}

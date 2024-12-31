package model

type Lineup struct {
	ID              LineupID
	Team            Team
	StartingMembers []*StartingMember
	Game            *Game //DHの有無をゲームから参照するため
}

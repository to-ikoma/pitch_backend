package model

type Lineup struct {
	ID              LineupID
	Team            Team
	StartingMembers []*StartingMember
}

package model

type Lineup struct {
	ID              LineupID
	Team            *Team
	StartingMembers []*StartingMember
}

// const (
// 	NotEnableDHMembers = 9
// 	EnableDHMembers    = 10
// )

// type OptionalLinup struct {
// 	Team            *Team
// 	StartingMembers []*StartingMember
// }

// func NewLineup(option *OptionalLinup) *Lineup {

// 	// startingMembers := NewStartingMembers()
// 	// if isEnableDH {
// 	// 	for i := 0; i < NotEnableDHMembers; i++ {
// 	// 		startingMembers = NewStartingMemberWithEnableDH()
// 	// 	}
// 	// }
// 	return &Lineup{
// 		ID:              NewLineupID(),
// 		Team:            option.Team,
// 		StartingMembers: option.StartingMembers,
// 	}
// }

package controller

import (
	"pitch_backend/domain/model"
	pb "pitch_backend/presentation/pb/api/pitchbackend/v1"
)

func ToGameEntity(pbGame *pb.Game) *model.Game {

	var date *model.Date
	if pbGame.Date != nil {
		date.Year = pbGame.Date.Year
		date.Month = pbGame.Date.Month
		date.Day = pbGame.Date.Day
	}

	return &model.Game{
		ID:               model.NewGameID(),
		Date:             date,
		Title:            pbGame.Title,
		IsEnableDH:       pbGame.DhEnbabled,
		FirstTeamLineup:  toLineupEntity(pbGame.FirstTeam),
		SecondTeamLineup: toLineupEntity(pbGame.SecondTeam),
	}
}

func toLineupEntity(pbLineup *pb.Lineup) *model.Lineup {

	return &model.Lineup{
		ID:              model.NewLineupID(),
		Team:            toTeamEntity(pbLineup.Team),
		StartingMembers: toStartingMembersEntity(pbLineup.StartingMembers),
	}

}

func toTeamEntity(pbTeam *pb.Team) *model.Team {
	return &model.Team{
		ID:      model.NewTeamID(),
		Name:    pbTeam.Name,
		Players: toPlayersEntity(pbTeam.Players),
	}
}

func toPlayerEntity(pbPlayer *pb.Player, position pb.Position) *model.Player {
	positions := []model.Position{model.Position(position)}
	return &model.Player{
		ID:                model.NewPlayerID(),
		Name:              pbPlayer.Name,
		FavoritePositions: positions,
	}
}

func toPlayersEntity(pbPlayers []*pb.Player) []*model.Player {
	return nil
}

func toStartingMembersEntity(pbStartingMembers []*pb.StartingMember) []*model.StartingMember {
	var startingMembers []*model.StartingMember

	for _, member := range pbStartingMembers {
		startingMembers = append(startingMembers,
			&model.StartingMember{
				ID:          model.NewStartingMemberID(),
				OrderNumber: int(member.OrderNumber),
				Player:      toPlayerEntity(member.Player, member.Position),
				Position:    model.Position(member.Position),
			},
		)
	}

	return startingMembers
}

package controller

import (
	"context"
	"pitch_backend/usecase"

	pb "pitch_backend/presentation/pb/api/pitchbackend/v1"
	svc "pitch_backend/presentation/pb/api/pitchbackend/v1/pitchbackendv1connect"

	"connectrpc.com/connect"
	"github.com/samber/do"
)

type GameController struct {
	GameUsecase usecase.GameUsecase
}

func NewGameController(i *do.Injector) (svc.GameServiceHandler, error) {
	gameUsecase := do.MustInvoke[usecase.GameUsecase](i)

	return &GameController{
		GameUsecase: gameUsecase,
	}, nil
}

func (c *GameController) CreateGame(ctx context.Context, req *connect.Request[pb.CreateGameRequest]) (*connect.Response[pb.CreateGameResponse], error) {
	game := ToGameEntity(req.Msg.Game)
	res := connect.NewResponse(&pb.CreateGameResponse{})
	err := c.GameUsecase.Create(ctx, game)
	if err != nil {
		return res, err
	}
	return res, nil
}

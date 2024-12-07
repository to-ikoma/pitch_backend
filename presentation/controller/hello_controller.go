package controller

import (
	"context"

	pb "pitch_backend/presentation/pb/api/pitchbackend/v1"
	svc "pitch_backend/presentation/pb/api/pitchbackend/v1/pitchbackendv1connect"
	"pitch_backend/usecase"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"github.com/samber/do"
)

type helloController struct {
	helloUsecase usecase.HelloUsecase
}

func NewHelloController(i *do.Injector) (svc.HelloServiceHandler, error) {
	helloUsecase := do.MustInvoke[usecase.HelloUsecase](i)

	return &helloController{
		helloUsecase: helloUsecase,
	}, nil
}

func (c *helloController) GetHello(ctx context.Context, req *connect.Request[pb.GetHelloRequest]) (*connect.Response[pb.GetHelloResponse], error) {
	helloModel, err := c.helloUsecase.Get()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Hello")
	}

	pbHello, err := FromHelloEntity(helloModel)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Hello")
	}

	return &connect.Response[pb.GetHelloResponse]{
		Msg: pbHello,
	}, nil

}

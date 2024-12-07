package controller

import (
	"errors"
	"pitch_backend/domain/model"
	pb "pitch_backend/presentation/pb/api/pitchbackend/v1"
)

func FromHelloEntity(entity *model.Hello) (res *pb.GetHelloResponse, err error) {
	if entity == nil {
		return nil, errors.New("hello entity is null")
	}

	return &pb.GetHelloResponse{
		Message: entity.Message,
	}, nil
}

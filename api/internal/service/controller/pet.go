package controller

import (
	"context"
	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/pet"
	"golang-grpc-recap/internal/service/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type PetHandler struct {
	pb.UnimplementedPetServiceServer
	UsecaseReceiver usecase.PetReceiver
}

func (h PetHandler) GetMyPet(ctx context.Context, request *emptypb.Empty) (*pb.GetPetReply, error) {
	output, err := h.UsecaseReceiver.GetMyPet(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetPetReply{
		Name:    output.Name,
		Species: output.Species,
		Breed:   output.Breed,
		Color:   output.Color,
		Age:     output.Age,
	}, nil
}

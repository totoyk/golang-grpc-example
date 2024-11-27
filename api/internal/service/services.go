package service

import (
	"google.golang.org/grpc"

	"golang-grpc-recap/internal/infra/repository"
	"golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/helloworld"
	"golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/pet"
	"golang-grpc-recap/internal/service/controller"
	"golang-grpc-recap/internal/service/usecase"
)

func RegisterServices(s *grpc.Server) {
	helloworld.RegisterGreeterServer(s, &controller.HelloworldHandler{
		UsecaseReceiver: usecase.NewHelloworldInteractor(
			repository.NewGreetRepository(),
		),
	})
	pet.RegisterPetServiceServer(s, &controller.PetHandler{
		UsecaseReceiver: usecase.NewPetInteractor(
			repository.NewPetRepository(),
		),
	})
}

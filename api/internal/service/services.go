package service

import (
	"google.golang.org/grpc"

	"golang-grpc-recap/internal/infra/repository"
	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/helloworld"
	"golang-grpc-recap/internal/service/controller"
	"golang-grpc-recap/internal/service/usecase"
)

func RegisterServices(s *grpc.Server) {
	pb.RegisterGreeterServer(s, &controller.HelloworldHandler{
		UsecaseReceiver: usecase.NewHelloworldInteractor(
			repository.NewGreetRepository(),
		),
	})
}

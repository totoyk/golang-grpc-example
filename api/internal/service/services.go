package service

import (
	"google.golang.org/grpc"

	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/helloworld"
	"golang-grpc-recap/internal/service/controller"
)

func RegisterServices(s *grpc.Server) {
	pb.RegisterGreeterServer(s, &controller.HelloworldHandler{})
}

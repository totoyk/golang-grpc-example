package controller

import (
	"context"

	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/helloworld"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HelloworldHandler struct {
	pb.UnimplementedGreeterServer
}

func (h HelloworldHandler) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello! " + request.Name}, nil
}
func (h HelloworldHandler) SayRepeatHello(ctx *pb.RepeatHelloRequest, srv pb.Greeter_SayRepeatHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayRepeatHello not implemented")
}

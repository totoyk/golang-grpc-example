package controller

import (
	"context"

	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/helloworld"
	"golang-grpc-recap/internal/service/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HelloworldHandler struct {
	pb.UnimplementedGreeterServer
	UsecaseReceiver usecase.HelloworldReceiver
}

func (h HelloworldHandler) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	output, err := h.UsecaseReceiver.SayHello(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: output}, nil
}
func (h HelloworldHandler) SayRepeatHello(ctx *pb.RepeatHelloRequest, srv pb.Greeter_SayRepeatHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayRepeatHello not implemented")
}

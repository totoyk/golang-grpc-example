package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "golang-grpc-recap/pb/github.com/totoyk/golang-grpc-recap/proto/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
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

func Run() int {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return 1
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &HelloworldHandler{})
	reflection.Register(grpcServer)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	grpcServer.GracefulStop()
	return 0
}

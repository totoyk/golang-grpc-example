package internal

import (
	"fmt"
	"golang-grpc-recap/internal/service"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() int {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return 1
	}

	grpcServer := grpc.NewServer()
	service.RegisterServices(grpcServer)
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

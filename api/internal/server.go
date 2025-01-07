package internal

import (
	"fmt"
	"golang-grpc-recap/internal/service"
	"golang-grpc-recap/internal/utils"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/providers/logrus/v2"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() int {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return 1
	}

	// TODO: enable log level from env
	logger, err := utils.NewLogger("info")

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_logging.UnaryServerInterceptor(grpc_logrus.InterceptorLogger(logger)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logging.StreamServerInterceptor(grpc_logrus.InterceptorLogger(logger)),
		),
	)
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

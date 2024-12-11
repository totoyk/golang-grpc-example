package controller

import (
	"context"
	repository "golang-grpc-recap/internal/infra/repository"
	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/helloworld"
	services "golang-grpc-recap/internal/service/controller"
	usecase "golang-grpc-recap/internal/service/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloworldService(t *testing.T) {
	// Init HelloworldHandler
	helloworldHandler := services.HelloworldHandler{
		UsecaseReceiver: usecase.NewHelloworldInteractor(
			repository.NewGreetRepository(),
		),
	}

	var res *pb.HelloReply
	var err error

	// Test SayHello
	res, err = helloworldHandler.SayHello(context.Background(), &pb.HelloRequest{Name: "TesterMan"})
	require.NoError(t, err)
	assert.Equal(t, "Hello, TesterMan", res.Message, "SayHello should return 'Hello TesterMan'")

	res, err = helloworldHandler.SayHello(context.Background(), &pb.HelloRequest{Name: "TesterWoman"})
	require.NoError(t, err)
	assert.NotEqual(t, "Hello, TesterMan", res.Message, "SayHello should not return 'Hello TesterMan'")
}

package usecase

import (
	"context"
	"golang-grpc-recap/internal/service/repository"
)

type HelloworldReceiver interface {
	SayHello(ctx context.Context, name string) (string, error)
}

type HelloworldInteractor struct {
	GreetRepository repository.GreetRepository
}

func NewHelloworldInteractor(greetRepository repository.GreetRepository) *HelloworldInteractor {
	return &HelloworldInteractor{
		GreetRepository: greetRepository,
	}
}

func (h *HelloworldInteractor) SayHello(ctx context.Context, name string) (string, error) {
	return h.GreetRepository.FindGreet() + name, nil
}

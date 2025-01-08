package usecase

import (
	"context"
	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/pet"
	"golang-grpc-recap/internal/service/repository"
	"golang-grpc-recap/internal/utils/zerrors"
)

type PetReceiver interface {
	GetMyPet(ctx context.Context) (*pb.Pet, error)
}

type PetInteractor struct {
	PetRepository repository.PetRepository
}

func NewPetInteractor(petRepository repository.PetRepository) *PetInteractor {
	return &PetInteractor{
		PetRepository: petRepository,
	}
}

func (p *PetInteractor) GetMyPet(ctx context.Context) (*pb.Pet, error) {
	pet, err := p.PetRepository.FindByOwner("owner")
	if err := zerrors.As(err); err != nil {
		return nil, err
	}
	return pet, nil
}

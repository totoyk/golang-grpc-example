package repository

import (
	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/pet"
)

type PetRepository struct{}

func NewPetRepository() *PetRepository {
	return &PetRepository{}
}

func (p *PetRepository) FindByOwner(owner string) (*pb.Pet, error) {
	return &pb.Pet{
		Name:    "Example Pet",
		Species: "Dog",
		Breed:   "Poodle",
		Color:   "White",
		Age:     "3",
	}, nil
}

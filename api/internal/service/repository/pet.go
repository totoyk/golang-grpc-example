package repository

import (
	pb "golang-grpc-recap/internal/pb/github.com/totoyk/golang-grpc-recap/proto/pet"
)

type PetRepository interface {
	FindByOwner(owner string) (*pb.Pet, error)
}

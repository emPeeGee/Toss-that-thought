package service

import (
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/emPeeee/ttt/pkg/repository"
)

type Service struct {
	Authorization
	Thought
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.Authorization),
		Thought:       NewThoughtService(repos.Thought),
	}
}

type Authorization interface {
}

type Thought interface {
	Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error)
	RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error)
	RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error)
	IsThoughtValid(thoughtKey string) (bool, error)
	CheckMetadataExists(metadataKey string) (bool, error)
	BurnThought(metadataKey, passphrase string) (bool, error)
}

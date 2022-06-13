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
	Create(entity.ThoughtInput) (entity.ThoughtCreateResponse, error)
	Metadata(metadataKey string) (entity.ThoughtMetadataResponse, error)
	CheckThoughtExists(thoughtKey string) (bool, error)
	AccessThought(thoughtKey, passphrase string) (entity.AccessThoughtResponse, error)
}

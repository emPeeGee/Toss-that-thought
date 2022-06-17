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
	RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtPassphraseInput, error)
	CheckThoughtExists(thoughtKey string) (bool, error)
	BurnThought(thoughtKey, passphrase string) (bool, error)
}

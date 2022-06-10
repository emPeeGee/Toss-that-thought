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
	Test() (interface{}, error)
	Create(entity.ThoughtInput) (entity.ThoughtResponse, error)
}

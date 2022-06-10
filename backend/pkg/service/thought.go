package service

import (
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/emPeeee/ttt/pkg/repository"
)

type ThoughtService struct {
	repo repository.Thought
}

func NewThoughtService(repo repository.Thought) *ThoughtService {
	return &ThoughtService{repo: repo}
}

func (s *ThoughtService) Test() (interface{}, error) {
	return s.repo.Test()
}

func (s *ThoughtService) Create(input entity.ThoughtInput) (entity.ThoughtResponse, error) {
	return s.repo.Create(input)
}

func (s *ThoughtService) Metadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {
	return s.repo.Metadata(metadataKey)
}

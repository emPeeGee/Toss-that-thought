package service

import "github.com/emPeeee/ttt/pkg/repository"

type ThoughtService struct {
	repo repository.Thought
}

func NewThoughtService(repo repository.Thought) *ThoughtService {
	return &ThoughtService{repo: repo}
}

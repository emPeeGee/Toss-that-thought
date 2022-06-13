package service

import (
	"errors"
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/emPeeee/ttt/pkg/repository"
	"golang.org/x/crypto/bcrypt"
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

func (s *ThoughtService) Create(input entity.ThoughtInput) (entity.ThoughtCreateResponse, error) {
	hashedPassphrase, err := HashPassphrase(input.Passphrase)

	if err != nil {
		return entity.ThoughtCreateResponse{}, err
	}

	input.Passphrase = hashedPassphrase
	return s.repo.Create(input)
}

func (s *ThoughtService) Metadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {
	return s.repo.Metadata(metadataKey)
}

func (s *ThoughtService) CheckThoughtExists(thoughtKey string) (bool, error) {
	return s.repo.CheckThoughtExists(thoughtKey)
}

func (s *ThoughtService) AccessThought(thoughtKey, passphrase string) (entity.AccessThoughtResponse, error) {
	hashedPassphrase, err := s.repo.GetPassphraseOfThought(thoughtKey)
	if err != nil {
		return entity.AccessThoughtResponse{}, err
	}

	if CheckPasswordHashes(passphrase, hashedPassphrase) == false {
		return entity.AccessThoughtResponse{}, errors.New("password does not match")
	}

	return s.repo.AccessThought(thoughtKey, hashedPassphrase)
}

func HashPassphrase(passphrase string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passphrase), 14)
	return string(bytes), err
}

func CheckPasswordHashes(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

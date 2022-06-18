package service

import (
	"errors"
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/emPeeee/ttt/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ThoughtService struct {
	repo repository.Thought
}

func NewThoughtService(repo repository.Thought) *ThoughtService {
	return &ThoughtService{repo: repo}
}

func (s *ThoughtService) Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error) {
	if len(input.Passphrase) != 0 {
		hashedPassphrase, err := HashPassphrase(input.Passphrase)
		if err != nil {
			return entity.ThoughtCreateResponse{}, err
		}
		input.Passphrase = hashedPassphrase
	}

	return s.repo.Create(input)
}

func (s *ThoughtService) RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {
	response, err := s.repo.RetrieveMetadata(metadataKey)
	response.AbbreviatedThoughtKey = response.AbbreviatedThoughtKey[:6]
	return response, err
}

func (s *ThoughtService) IsThoughtValid(thoughtKey string) (bool, error) {
	thoughtValidityInfo, err := s.repo.RetrieveThoughtValidity(thoughtKey)
	if err != nil {
		return false, err
	}

	now := time.Now()

	// if lifetime is passed, or is burned or is viewed return err
	if now.After(thoughtValidityInfo.Lifetime) || thoughtValidityInfo.IsBurned || thoughtValidityInfo.IsViewed {
		return false, errors.New("it either never existed or already has been viewed")
	}

	return true, nil
}

func (s *ThoughtService) CheckMetadataExists(metadataKey string) (bool, error) {
	return s.repo.CheckMetadataExists(metadataKey)
}

func (s *ThoughtService) RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error) {
	hashedPassphrase, err := s.repo.GetPassphraseOfThoughtByThoughtKey(thoughtKey)
	if err != nil {
		return entity.ThoughtResponse{}, err
	}

	if len(hashedPassphrase) != 0 && CheckPasswordHashes(passphrase, hashedPassphrase) == false {
		return entity.ThoughtResponse{}, errors.New("password does not match")
	}

	thoughtResponse, err := s.repo.RetrieveThought(thoughtKey, hashedPassphrase)
	if err != nil {
		return entity.ThoughtResponse{}, err
	}

	if err := s.repo.MarkAsViewed(thoughtKey, hashedPassphrase); err != nil {
		return entity.ThoughtResponse{}, errors.New("something went wrong")
	}

	return thoughtResponse, nil
}

func (s *ThoughtService) BurnThought(metadataKey, passphrase string) (bool, error) {
	thoughtMetadata, err := s.repo.RetrieveMetadata(metadataKey)
	if err != nil {
		return false, err
	}

	if thoughtMetadata.IsBurned {
		return false, errors.New("thought is already burned")
	}

	hashedPassphrase, err := s.repo.GetPassphraseOfThoughtByMetadataKey(metadataKey)
	if err != nil {
		return false, err
	}

	if len(hashedPassphrase) != 0 && CheckPasswordHashes(passphrase, hashedPassphrase) == false {
		return false, errors.New("password does not match")
	}

	return s.repo.BurnThought(metadataKey, hashedPassphrase)
}

func HashPassphrase(passphrase string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passphrase), 14)
	return string(bytes), err
}

func CheckPasswordHashes(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

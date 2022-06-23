package thought

import (
	"errors"
	"github.com/emPeeee/ttt/pkg/crypt"
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"time"
)

type Service interface {
	Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error)
	RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error)
	RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error)
	IsThoughtValid(thoughtKey string) (bool, error)
	CheckMetadataExists(metadataKey string) (bool, error)
	BurnThought(metadataKey, passphrase string) (bool, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewThoughtService(repo Repository, logger log.Logger) *service {
	return &service{repo: repo, logger: logger}
}

func (s *service) Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error) {
	if len(input.Passphrase) != 0 {
		hashedPassphrase, err := crypt.HashPassphrase(input.Passphrase)
		if err != nil {
			return entity.ThoughtCreateResponse{}, err
		}
		input.Passphrase = hashedPassphrase
	}

	createdThought, err := s.repo.Create(input)
	if err == nil {
		createdThought.AbbreviatedThoughtKey = createdThought.ThoughtKey[:6]
	}

	return createdThought, err
}

func (s *service) RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {
	metadata, err := s.repo.RetrieveMetadata(metadataKey)
	if err == nil {
		metadata.AbbreviatedThoughtKey = metadata.AbbreviatedThoughtKey[:6]
	}

	if time.Now().After(metadata.Lifetime) {
		return entity.ThoughtMetadataResponse{}, errors.New("it either never existed or already has been viewed")
	}

	return metadata, err
}

func (s *service) IsThoughtValid(thoughtKey string) (bool, error) {
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

func (s *service) CheckMetadataExists(metadataKey string) (bool, error) {
	return s.repo.CheckMetadataExists(metadataKey)
}

func (s *service) RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error) {
	hashedPassphrase, err := s.repo.GetPassphraseOfThoughtByThoughtKey(thoughtKey)
	if err != nil {
		return entity.ThoughtResponse{}, err
	}

	if len(hashedPassphrase) != 0 && crypt.CheckPasswordHashes(passphrase, hashedPassphrase) == false {
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

func (s *service) BurnThought(metadataKey, passphrase string) (bool, error) {
	thoughtMetadata, err := s.repo.RetrieveMetadata(metadataKey)
	if err != nil {
		return false, err
	}

	if time.Now().After(thoughtMetadata.Lifetime) {
		return false, errors.New("it either never existed or already has been viewed")
	}

	if thoughtMetadata.IsBurned {
		return false, errors.New("thought is already burned")
	}

	if thoughtMetadata.IsViewed {
		return false, errors.New("thought is already viewed")
	}

	if time.Now().After(thoughtMetadata.Lifetime) {
		return false, errors.New("cannot burn, lifetime passed")
	}

	hashedPassphrase, err := s.repo.GetPassphraseOfThoughtByMetadataKey(metadataKey)
	if err != nil {
		return false, err
	}

	if len(hashedPassphrase) != 0 && crypt.CheckPasswordHashes(passphrase, hashedPassphrase) == false {
		return false, errors.New("password does not match")
	}

	return s.repo.BurnThought(metadataKey, hashedPassphrase)
}

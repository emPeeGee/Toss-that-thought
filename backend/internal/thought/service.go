package thought

import (
	"errors"
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/crypt"
	"github.com/emPeeee/ttt/pkg/log"
	"time"
)

type Service interface {
	Create(input CreateDTO, userId *uint) (CreateResponse, error)
	RetrieveMetadata(metadataKey string) (MetadataResponse, error)
	RetrieveThoughtByPassphrase(thoughtKey, passphrase string) (ViewThoughtResponse, error)
	RetrieveThoughtPassphraseInfo(thoughtKey string) (PassphraseInformationResponse, error)
	IsThoughtValid(thoughtKey string) error
	CheckMetadataExists(metadataKey string) error
	BurnThought(metadataKey, passphrase string) error
	GetThoughtsMetadataByUser(userId uint) ([]MetadataResponse, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewThoughtService(repo Repository, logger log.Logger) *service {
	return &service{repo: repo, logger: logger}
}

func (s *service) Create(input CreateDTO, userId *uint) (CreateResponse, error) {
	thought := entity.Thought{
		Thought:  input.Thought,
		Lifetime: input.Lifetime,
		UserID:   userId,
	}

	// If passphrase is empty, when omit hashing
	if len(input.Passphrase) != 0 {
		err := thought.HashPassphrase(input.Passphrase)
		if err != nil {
			return CreateResponse{}, err
		}
	}

	createdThought, err := s.repo.Create(thought)
	if err == nil {
		createdThought.AbbreviatedThoughtKey = createdThought.ThoughtKey.String()[:6]
	}

	return createdThought, err
}

func (s *service) RetrieveMetadata(metadataKey string) (MetadataResponse, error) {
	metadata, err := s.repo.RetrieveMetadata(metadataKey)
	if err == nil {
		metadata.AbbreviatedThoughtKey = metadata.AbbreviatedThoughtKey[:6]
	}

	if time.Now().After(metadata.Lifetime) {
		return MetadataResponse{}, errors.New("it either never existed or already has been viewed")
	}

	return metadata, err
}

func (s *service) IsThoughtValid(thoughtKey string) error {
	thoughtValidityInfo, err := s.repo.RetrieveThoughtValidity(thoughtKey)
	if err != nil {
		return err
	}

	now := time.Now()

	// if lifetime is passed, or is burned or is viewed return err
	if now.After(thoughtValidityInfo.Lifetime) || thoughtValidityInfo.IsBurned || thoughtValidityInfo.IsViewed {
		return errors.New("it either never existed or already has been viewed")
	}

	return nil
}

func (s *service) RetrieveThoughtPassphraseInfo(thoughtKey string) (PassphraseInformationResponse, error) {
	err := s.IsThoughtValid(thoughtKey)
	if err != nil {
		return PassphraseInformationResponse{}, err
	}

	hashedPassphrase, err := s.repo.GetPassphraseOfThoughtByThoughtKey(thoughtKey)
	if err != nil {
		return PassphraseInformationResponse{}, err
	}

	canPassphraseSkipped := len(hashedPassphrase) <= 0

	return PassphraseInformationResponse{CanPassphraseBeSkipped: canPassphraseSkipped}, nil
}

func (s *service) CheckMetadataExists(metadataKey string) error {
	return s.repo.CheckMetadataExists(metadataKey)
}

func (s *service) RetrieveThoughtByPassphrase(thoughtKey, passphrase string) (ViewThoughtResponse, error) {
	hashedPassphrase, err := s.repo.GetPassphraseOfThoughtByThoughtKey(thoughtKey)
	if err != nil {
		return ViewThoughtResponse{}, err
	}

	if len(hashedPassphrase) != 0 && crypt.CheckPasswordHashes(passphrase, hashedPassphrase) != nil {
		return ViewThoughtResponse{}, errors.New("password does not match")
	}

	thoughtResponse, err := s.repo.RetrieveThoughtByPassphrase(thoughtKey, hashedPassphrase)
	if err != nil {
		return ViewThoughtResponse{}, err
	}

	if err := s.repo.MarkAsViewed(thoughtKey, hashedPassphrase); err != nil {
		return ViewThoughtResponse{}, errors.New("something went wrong")
	}

	return thoughtResponse, nil
}

func (s *service) BurnThought(metadataKey, passphrase string) error {
	thoughtMetadata, err := s.repo.RetrieveMetadata(metadataKey)
	if err != nil {
		return err
	}

	if time.Now().After(thoughtMetadata.Lifetime) {
		return errors.New("it either never existed or already has been viewed")
	}

	if thoughtMetadata.IsBurned {
		return errors.New("thought is already burned")
	}

	if thoughtMetadata.IsViewed {
		return errors.New("thought is already viewed")
	}

	if time.Now().After(thoughtMetadata.Lifetime) {
		return errors.New("cannot burn, lifetime passed")
	}

	hashedPassphrase, err := s.repo.GetPassphraseOfThoughtByMetadataKey(metadataKey)
	if err != nil {
		return err
	}

	if len(hashedPassphrase) != 0 && crypt.CheckPasswordHashes(passphrase, hashedPassphrase) != nil {
		return errors.New("password does not match")
	}

	return s.repo.BurnThought(metadataKey, hashedPassphrase)
}

func (s *service) GetThoughtsMetadataByUser(userId uint) ([]MetadataResponse, error) {
	thoughtsMetadata, err := s.repo.GetThoughtsMetadataByUser(userId)
	if err != nil {
		return nil, err
	}

	for idx, metadata := range thoughtsMetadata {
		thoughtsMetadata[idx].AbbreviatedThoughtKey = metadata.AbbreviatedThoughtKey[:6]
	}

	return thoughtsMetadata, nil
}

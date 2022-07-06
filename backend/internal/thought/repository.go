package thought

import (
	"errors"
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	Create(input entity.Thought) (CreateResponse, error)
	RetrieveMetadata(metadataKey string) (MetadataResponse, error)
	RetrieveThoughtByPassphrase(thoughtKey, passphrase string) (ViewThoughtResponse, error)
	RetrieveThoughtValidity(thoughtKey string) (ValidityInformation, error)
	CheckMetadataExists(metadataKey string) error
	BurnThought(metadataKey, passphrase string) error
	MarkAsViewed(thoughtKey, passphrase string) error
	GetPassphraseOfThoughtByMetadataKey(metadataKey string) (string, error)
	GetPassphraseOfThoughtByThoughtKey(thoughtKey string) (string, error)
	GetThoughtsMetadataByUser(userId uint) ([]MetadataResponse, error)
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *repository {
	return &repository{db: db, logger: logger}
}

func (r *repository) Create(thought entity.Thought) (CreateResponse, error) {

	if err := r.db.Create(&thought).Error; err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{
		Lifetime:              thought.Lifetime,
		MetadataKey:           thought.MetadataKey,
		ThoughtKey:            thought.ThoughtKey,
		AbbreviatedThoughtKey: thought.ThoughtKey.String(),
		IsBurned:              thought.IsBurned,
	}, nil
}

func (r *repository) RetrieveMetadata(metadataKey string) (MetadataResponse, error) {

	var thoughtMetadata MetadataResponse
	err := r.db.Model(&entity.Thought{}).First(&thoughtMetadata, "metadata_key = ?", metadataKey).Error

	return thoughtMetadata, err
}

func (r *repository) RetrieveThoughtValidity(thoughtKey string) (ValidityInformation, error) {
	var thoughtValidityInfo ValidityInformation

	if err := r.db.Model(&entity.Thought{}).First(&thoughtValidityInfo, "thought_key = ?", thoughtKey).Error; err != nil {
		return ValidityInformation{}, err
	}

	return thoughtValidityInfo, nil
}

func (r *repository) CheckMetadataExists(metadataKey string) error {
	var exists bool

	if err := r.db.Model(&entity.Thought{}).Select("count(*) > 0").Where("metadata_key = ?", metadataKey).Find(&exists).Error; err != nil {
		return err
	}

	// TODO: Like that
	if !exists {
		return errors.New("such thought does not exist")
	}

	return nil
}

func (r *repository) RetrieveThoughtByPassphrase(thoughtKey, passphrase string) (ViewThoughtResponse, error) {
	var thoughtResponse ViewThoughtResponse

	if err := r.db.Model(&entity.Thought{}).First(&thoughtResponse, "thought_key = ? AND passphrase = ?", thoughtKey, passphrase).Error; err != nil {
		return ViewThoughtResponse{}, err
	}

	return thoughtResponse, nil
}

func (r *repository) MarkAsViewed(thoughtKey, passphrase string) error {
	now := time.Now()

	result := r.db.Model(&entity.Thought{}).Where("passphrase = ? AND thought_key = ?", passphrase, thoughtKey).Updates(entity.Thought{
		IsViewed:   true,
		ViewedDate: &now,
	})

	if result.Error != nil && result.RowsAffected <= 0 {
		return result.Error
	}

	return nil
}

func (r *repository) BurnThought(metadataKey, passphrase string) error {
	now := time.Now()

	result := r.db.Model(&entity.Thought{}).Where("passphrase = ? AND metadata_key = ?", passphrase, metadataKey).Updates(&entity.Thought{
		IsBurned:   true,
		BurnedDate: &now,
	})

	if result.Error != nil && result.RowsAffected <= 0 {
		return result.Error
	}

	return nil
}

// to lower case this?
func (r *repository) GetPassphraseOfThoughtByMetadataKey(metadataKey string) (string, error) {
	var passphrase string

	if err := r.db.Model(&entity.Thought{}).Select("passphrase").Where("metadata_key = ?", metadataKey).First(&passphrase).Error; err != nil {
		return "", err
	}

	return passphrase, nil
}

func (r *repository) GetPassphraseOfThoughtByThoughtKey(thoughtKey string) (string, error) {
	var passphrase string

	if err := r.db.Model(&entity.Thought{}).Select("passphrase").Where("thought_key = ?", thoughtKey).First(&passphrase).Error; err != nil {
		return "", err
	}

	return passphrase, nil
}

func (r *repository) GetThoughtsMetadataByUser(userId uint) ([]MetadataResponse, error) {
	var thoughtsMetadata []MetadataResponse
	var user entity.User

	if err := r.db.Model(&user).Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&user).Order("created_date asc").Association("Thoughts").Find(&thoughtsMetadata, "lifetime > current_timestamp"); err != nil {
		return nil, err
	}

	return thoughtsMetadata, nil
}

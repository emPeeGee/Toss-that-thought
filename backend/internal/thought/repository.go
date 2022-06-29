package thought

import (
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error)
	RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error)
	RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error)
	RetrieveThoughtValidity(thoughtKey string) (entity.ThoughtValidityInformation, error)
	CheckMetadataExists(metadataKey string) (bool, error)
	BurnThought(metadataKey, passphrase string) (bool, error)
	MarkAsViewed(thoughtKey, passphrase string) error
	GetPassphraseOfThoughtByMetadataKey(metadataKey string) (string, error)
	GetPassphraseOfThoughtByThoughtKey(thoughtKey string) (string, error)
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *repository {
	return &repository{db: db, logger: logger}
}

func (r *repository) Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error) {

	thought := entity.Thought{
		Thought:    input.Thought,
		Passphrase: input.Passphrase,
		Lifetime:   input.Lifetime,
	}

	if err := r.db.Create(&thought).Error; err != nil {
		return entity.ThoughtCreateResponse{}, err
	}

	return entity.ThoughtCreateResponse{
		Lifetime:              thought.Lifetime,
		MetadataKey:           thought.MetadataKey,
		ThoughtKey:            thought.ThoughtKey,
		AbbreviatedThoughtKey: thought.ThoughtKey.String(),
		IsBurned:              thought.IsBurned,
	}, nil
}

func (r *repository) RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {

	var thoughtMetadata entity.ThoughtMetadataResponse
	err := r.db.Model(&entity.Thought{}).First(&thoughtMetadata, "metadata_key = ?", metadataKey).Error

	return thoughtMetadata, err
}

func (r *repository) RetrieveThoughtValidity(thoughtKey string) (entity.ThoughtValidityInformation, error) {
	var thoughtValidityInfo entity.ThoughtValidityInformation

	if err := r.db.Model(&entity.Thought{}).First(&thoughtValidityInfo, "thought_key = ?", thoughtKey).Error; err != nil {
		return entity.ThoughtValidityInformation{}, err
	}

	return thoughtValidityInfo, nil
}

func (r *repository) CheckMetadataExists(metadataKey string) (bool, error) {
	var exists bool

	if err := r.db.Model(&entity.Thought{}).Select("count(*) > 0").Where("metadata_key = ?", metadataKey).Find(&exists).Error; err != nil || !exists {
		return false, err
	}

	return true, nil
}

func (r *repository) RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error) {
	var thoughtResponse entity.ThoughtResponse

	if err := r.db.Model(&entity.Thought{}).First(&thoughtResponse, "thought_key = ? AND passphrase = ?", thoughtKey, passphrase).Error; err != nil {
		return entity.ThoughtResponse{}, err
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

func (r *repository) BurnThought(metadataKey, passphrase string) (bool, error) {
	now := time.Now()

	result := r.db.Model(&entity.Thought{}).Where("passphrase = ? AND metadata_key = ?", passphrase, metadataKey).Updates(&entity.Thought{
		IsBurned:   true,
		BurnedDate: &now,
	})

	if result.Error != nil && result.RowsAffected <= 0 {
		return false, result.Error
	}

	return true, nil
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

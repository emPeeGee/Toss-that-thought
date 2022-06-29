package thought

import (
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/jmoiron/sqlx"
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
	db     *sqlx.DB
	gorm   *gorm.DB
	logger log.Logger
}

func NewRepository(db *sqlx.DB, gorm *gorm.DB, logger log.Logger) *repository {
	return &repository{db: db, gorm: gorm, logger: logger}
}

func (r *repository) Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error) {
	//var thoughtResponse entity.ThoughtCreateResponse
	//createThoughtQuery := `INSERT INTO thoughts(thought, passphrase, lifetime) VALUES ($1, $2, $3) RETURNING metadata_key, thought_key, is_burned, lifetime`
	//row := r.db.QueryRowx(createThoughtQuery, input.Thought, input.Passphrase, input.Lifetime)
	//
	//if err := row.StructScan(&thoughtResponse); err != nil {
	//	return thoughtResponse, err
	//}

	thought := entity.Thought{
		Thought:    input.Thought,
		Passphrase: input.Passphrase,
		Lifetime:   input.Lifetime,
		//MetadataKey: uuid.NewString(),
		//ThoughtKey:  uuid.NewString(),
	}

	if err := r.gorm.Create(&thought).Error; err != nil {
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
	//var thoughtMetadata entity.ThoughtMetadataResponse
	//thoughtMetadataQuery := "SELECT th.lifetime, th.is_burned, th.burned_date, th.is_viewed, th.viewed_date, th.created_date, th.thought_key as abbreviated_thought_key FROM thoughts th WHERE th.metadata_key = $1"
	//err := r.db.Get(&thoughtMetadata, thoughtMetadataQuery, metadataKey)

	var thoughtMetadata entity.ThoughtMetadataResponse
	err := r.gorm.Model(&entity.Thought{}).First(&thoughtMetadata, "metadata_key = ?", metadataKey).Error

	return thoughtMetadata, err
}

func (r *repository) RetrieveThoughtValidity(thoughtKey string) (entity.ThoughtValidityInformation, error) {
	var thoughtValidityInfo entity.ThoughtValidityInformation
	//query := "SELECT th.thought_key, th.lifetime, th.is_burned, is_viewed FROM thoughts th WHERE th.thought_key = $1;"
	//err := r.db.Get(&thoughtValidityInfo, query, thoughtKey)
	//
	//if err != nil {
	//	return entity.ThoughtValidityInformation{}, err
	//}

	if err := r.gorm.Model(&entity.Thought{}).First(&thoughtValidityInfo, "thought_key = ?", thoughtKey).Error; err != nil {
		return entity.ThoughtValidityInformation{}, err
	}

	return thoughtValidityInfo, nil
}

func (r *repository) CheckMetadataExists(metadataKey string) (bool, error) {
	var exists bool
	//query := "SELECT exists(SELECT th.id FROM thoughts th WHERE th.metadata_key = $1);"
	//row := r.db.QueryRow(query, metadataKey)

	// Did not handle exists response
	// And made mistakes in which I incurc metadata with thoughtKeu because I did not check
	//err := row.Scan(&exists)
	//if err != nil || !exists {
	//	return false, errors.New("Row does not exist")
	//}

	if err := r.gorm.Model(&entity.Thought{}).Select("count(*) > 0").Where("metadata_key = ?", metadataKey).Find(&exists).Error; err != nil || !exists {
		return false, err
	}

	return true, nil
}

func (r *repository) RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error) {
	var thoughtResponse entity.ThoughtResponse
	//query := "SELECT th.thought from thoughts th WHERE th.thought_key = $1 AND th.passphrase = $2"
	//err := r.db.Get(&thoughtResponse, query, thoughtKey, passphrase)
	//
	//if err != nil {
	//	return entity.ThoughtResponse{}, err
	//}
	//
	//return thoughtResponse, nil

	if err := r.gorm.Model(&entity.Thought{}).First(&thoughtResponse, "thought_key = ? AND passphrase = ?", thoughtKey, passphrase).Error; err != nil {
		return entity.ThoughtResponse{}, err
	}

	return thoughtResponse, nil
}

func (r *repository) MarkAsViewed(thoughtKey, passphrase string) error {
	//query := "UPDATE thoughts SET is_viewed = true, viewed_date = current_timestamp WHERE thought_key = $1 AND passphrase = $2"
	//res, err := r.db.Exec(query, thoughtKey, passphrase)
	//rowsAffected, _ := res.RowsAffected()
	//
	//if err != nil || rowsAffected <= 0 {
	//	return errors.New("server error")
	//}

	now := time.Now()

	result := r.gorm.Model(&entity.Thought{}).Where("passphrase = ? AND thought_key = ?", passphrase, thoughtKey).Updates(entity.Thought{
		IsViewed:   true,
		ViewedDate: &now,
	})

	if result.Error != nil && result.RowsAffected <= 0 {
		return result.Error
	}

	return nil
}

func (r *repository) BurnThought(metadataKey, passphrase string) (bool, error) {
	//query := "UPDATE thoughts SET is_burned = true, burned_date = current_timestamp WHERE metadata_key = $1 AND passphrase = $2"
	//res, err := r.db.Exec(query, metadataKey, passphrase)
	//
	//if err != nil {
	//	return false, err
	//}
	//
	//nr, _ := res.RowsAffected()

	now := time.Now()

	result := r.gorm.Model(&entity.Thought{}).Where("passphrase = ? AND metadata_key = ?", passphrase, metadataKey).Updates(&entity.Thought{
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
	//query := "SELECT th.passphrase from thoughts th WHERE th.metadata_key = $1"
	//row := r.db.QueryRow(query, metadataKey)
	//
	//var hashedPassphrase string
	//err := row.Scan(&hashedPassphrase)
	//
	//if err != nil {
	//	return "", err
	//}

	var userPassphrase entity.UserPassphrase

	if err := r.gorm.Model(&entity.Thought{}).Where("metadata_key = ?", metadataKey).First(&userPassphrase).Error; err != nil {
		return "", err
	}

	return userPassphrase.Passphrase, nil
}

func (r *repository) GetPassphraseOfThoughtByThoughtKey(thoughtKey string) (string, error) {
	//query := "SELECT th.passphrase from thoughts th WHERE th.thought_key= $1"
	//row := r.db.QueryRow(query, thoughtKey)
	//
	//var hashedPassphrase string
	//err := row.Scan(&hashedPassphrase)
	//
	//if err != nil {
	//	return "", err
	//}
	//
	//return hashedPassphrase, nil

	var userPassphrase entity.UserPassphrase

	if err := r.gorm.Model(&entity.Thought{}).Where("thought_key = ?", thoughtKey).First(&userPassphrase).Error; err != nil {
		return "", err
	}

	return userPassphrase.Passphrase, nil
}

package thought

import (
	"errors"
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/jmoiron/sqlx"
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
	logger log.Logger
}

func NewRepository(db *sqlx.DB, logger log.Logger) *repository {
	return &repository{db: db, logger: logger}
}

func (r *repository) Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error) {
	var thoughtResponse entity.ThoughtCreateResponse
	createThoughtQuery := `INSERT INTO thoughts(thought, passphrase, lifetime) VALUES ($1, $2, $3) RETURNING metadata_key, thought_key, is_burned, lifetime`
	row := r.db.QueryRowx(createThoughtQuery, input.Thought, input.Passphrase, input.Lifetime)

	if err := row.StructScan(&thoughtResponse); err != nil {
		return thoughtResponse, err
	}

	return thoughtResponse, nil
}

func (r *repository) RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {
	var thoughtMetadata entity.ThoughtMetadataResponse

	thoughtMetadataQuery := "SELECT th.lifetime, th.is_burned, th.burned_date, th.is_viewed, th.viewed_date, th.created_date, th.thought_key as abbreviated_thought_key FROM thoughts th WHERE th.metadata_key = $1"
	err := r.db.Get(&thoughtMetadata, thoughtMetadataQuery, metadataKey)

	return thoughtMetadata, err
}

func (r *repository) RetrieveThoughtValidity(thoughtKey string) (entity.ThoughtValidityInformation, error) {
	var thoughtValidityInfo entity.ThoughtValidityInformation
	query := "SELECT th.thought_key, th.lifetime, th.is_burned, is_viewed FROM thoughts th WHERE th.thought_key = $1;"
	err := r.db.Get(&thoughtValidityInfo, query, thoughtKey)

	if err != nil {
		return entity.ThoughtValidityInformation{}, err
	}

	return thoughtValidityInfo, nil
}

func (r *repository) CheckMetadataExists(metadataKey string) (bool, error) {
	var exists bool
	query := "SELECT exists(SELECT th.id FROM thoughts th WHERE th.metadata_key = $1);"
	row := r.db.QueryRow(query, metadataKey)

	// Did not handle exists response
	// And made mistakes in which I incurc metadata with thoughtKeu because I did not check
	err := row.Scan(&exists)
	if err != nil || !exists {
		return false, errors.New("Row does not exist")
	}

	return true, nil
}

func (r *repository) RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error) {
	var thoughtResponse entity.ThoughtResponse
	query := "SELECT th.thought from thoughts th WHERE th.thought_key = $1 AND th.passphrase = $2"
	err := r.db.Get(&thoughtResponse, query, thoughtKey, passphrase)

	if err != nil {
		return entity.ThoughtResponse{}, err
	}

	return thoughtResponse, nil
}

func (r *repository) MarkAsViewed(thoughtKey, passphrase string) error {
	query := "UPDATE thoughts SET is_viewed = true, viewed_date = current_timestamp WHERE thought_key = $1 AND passphrase = $2"
	res, err := r.db.Exec(query, thoughtKey, passphrase)
	rowsAffected, _ := res.RowsAffected()

	if err != nil || rowsAffected <= 0 {
		return errors.New("server error")
	}

	return nil
}

func (r *repository) BurnThought(metadataKey, passphrase string) (bool, error) {
	query := "UPDATE thoughts SET is_burned = true, burned_date = current_timestamp WHERE metadata_key = $1 AND passphrase = $2"
	res, err := r.db.Exec(query, metadataKey, passphrase)

	if err != nil {
		return false, err
	}

	nr, _ := res.RowsAffected()
	r.logger.Info(nr)

	return true, nil
}

func (r *repository) GetPassphraseOfThoughtByMetadataKey(metadataKey string) (string, error) {
	query := "SELECT th.passphrase from thoughts th WHERE th.metadata_key = $1"
	row := r.db.QueryRow(query, metadataKey)

	var hashedPassphrase string
	err := row.Scan(&hashedPassphrase)

	if err != nil {
		return "", err
	}

	return hashedPassphrase, nil
}

func (r *repository) GetPassphraseOfThoughtByThoughtKey(thoughtKey string) (string, error) {
	query := "SELECT th.passphrase from thoughts th WHERE th.thought_key= $1"
	row := r.db.QueryRow(query, thoughtKey)

	var hashedPassphrase string
	err := row.Scan(&hashedPassphrase)

	if err != nil {
		return "", err
	}

	return hashedPassphrase, nil
}

package repository

import (
	"errors"
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ThoughtSql struct {
	db *sqlx.DB
}

func NewThoughtSql(db *sqlx.DB) *ThoughtSql {
	return &ThoughtSql{db: db}
}

// TODO: If there is no passphrase when what to send to front?
func (r *ThoughtSql) Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error) {
	var thoughtResponse entity.ThoughtCreateResponse
	createThoughtQuery := `INSERT INTO thoughts(thought, passphrase, lifetime) VALUES ($1, $2, $3) RETURNING metadata_key, thought_key, is_burned, lifetime`
	row := r.db.QueryRowx(createThoughtQuery, input.Thought, input.Passphrase, input.Lifetime)

	if err := row.StructScan(&thoughtResponse); err != nil {
		return thoughtResponse, err
	}

	return thoughtResponse, nil
}

func (r *ThoughtSql) RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {
	var thoughtMetadata entity.ThoughtMetadataResponse

	thoughtMetadataQuery := "SELECT th.lifetime, th.is_burned, th.created_date, th.thought_key as abbreviated_thought_key FROM thoughts th WHERE th.metadata_key = $1"
	err := r.db.Get(&thoughtMetadata, thoughtMetadataQuery, metadataKey)

	return thoughtMetadata, err
}

func (r *ThoughtSql) CheckThoughtExists(thoughtKey string) (bool, error) {
	var exists bool
	query := "SELECT exists(SELECT th.id FROM thoughts th WHERE th.thought_key = $1);"
	row := r.db.QueryRow(query, thoughtKey)

	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ThoughtSql) CheckMetadataExists(metadataKey string) (bool, error) {
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

func (r *ThoughtSql) RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error) {
	var thoughtResponse entity.ThoughtResponse
	query := "SELECT th.thought from thoughts th WHERE th.thought_key = $1 AND th.passphrase = $2"
	err := r.db.Get(&thoughtResponse, query, thoughtKey, passphrase)

	if err != nil {
		return entity.ThoughtResponse{}, err
	}

	return thoughtResponse, nil
}

func (r *ThoughtSql) BurnThought(metadataKey, passphrase string) (bool, error) {
	query := "UPDATE thoughts SET is_burned = true, burned_date = current_timestamp WHERE metadata_key = $1 AND passphrase = $2"
	res, err := r.db.Exec(query, metadataKey, passphrase)

	if err != nil {
		return false, err
	}

	nr, _ := res.RowsAffected()
	logrus.Info(nr)

	return true, nil
}

// TODO: Question, it is normal to create separate query for getting passoword hash? It is used just internally
func (r *ThoughtSql) GetPassphraseOfThoughtByMetadataKey(metadataKey string) (string, error) {
	query := "SELECT th.passphrase from thoughts th WHERE th.metadata_key = $1"
	row := r.db.QueryRow(query, metadataKey)

	var hashedPassphrase string
	err := row.Scan(&hashedPassphrase)

	if err != nil {
		return "", err
	}

	return hashedPassphrase, nil
}

func (r *ThoughtSql) GetPassphraseOfThoughtByThoughtKey(thoughtKey string) (string, error) {
	query := "SELECT th.passphrase from thoughts th WHERE th.thought_key= $1"
	row := r.db.QueryRow(query, thoughtKey)

	var hashedPassphrase string
	err := row.Scan(&hashedPassphrase)

	if err != nil {
		return "", err
	}

	return hashedPassphrase, nil
}

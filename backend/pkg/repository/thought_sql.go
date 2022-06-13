package repository

import (
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

func (r *ThoughtSql) Test() (interface{}, error) {
	var thoughts []entity.Thought
	err := r.db.Select(&thoughts, "SELECT * from thoughts th")
	logrus.Info(thoughts)

	return thoughts, err
}

func (r *ThoughtSql) Create(input entity.ThoughtInput) (entity.ThoughtCreateResponse, error) {
	var thoughtResponse entity.ThoughtCreateResponse
	createThoughtQuery := `INSERT INTO thoughts(thought, passphrase, lifetime) VALUES ($1, $2, $3) RETURNING metadata_key;`
	row := r.db.QueryRowx(createThoughtQuery, input.Thought, input.Passphrase, input.Lifetime)

	if err := row.StructScan(&thoughtResponse); err != nil {
		return thoughtResponse, err
	}

	return thoughtResponse, nil
}

func (r *ThoughtSql) Metadata(metadataKey string) (entity.ThoughtMetadataResponse, error) {
	var thoughtMetadata entity.ThoughtMetadataResponse

	thoughtMetadataQuery := "SELECT th.lifetime, th.is_burned, th.created_date, th.thought_key FROM thoughts th WHERE th.metadata_key = $1"
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

func (r *ThoughtSql) AccessThought(thoughtKey, passphrase string) (entity.AccessThoughtResponse, error) {
	var thoughtResponse entity.AccessThoughtResponse
	query := "SELECT th.thought from thoughts th WHERE th.thought_key = $1 AND th.passphrase = $2"
	err := r.db.Get(&thoughtResponse, query, thoughtKey, passphrase)

	if err != nil {
		return entity.AccessThoughtResponse{}, err
	}

	return thoughtResponse, nil
}

// TODO: Question, it is normal to create separate query for getting passoword hash? It is used just internally
func (r *ThoughtSql) GetPassphraseOfThought(thoughtKey string) (string, error) {
	query := "SELECT th.passphrase from thoughts th WHERE th.thought_key = $1"
	row := r.db.QueryRow(query, thoughtKey)

	var hashedPassphrase string
	err := row.Scan(&hashedPassphrase)
	logrus.Info(hashedPassphrase)

	if err != nil {
		return "", err
	}

	return hashedPassphrase, nil
}

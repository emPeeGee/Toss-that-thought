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

func (r *ThoughtSql) Create(input entity.ThoughtInput) (entity.ThoughtResponse, error) {
	var thoughtResponse entity.ThoughtResponse
	createThoughtQuery := `INSERT INTO thoughts(passphrase, lifetime) VALUES ($1, $2) RETURNING metadata_key;`
	row := r.db.QueryRowx(createThoughtQuery, input.Passphrase, input.Lifetime)

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

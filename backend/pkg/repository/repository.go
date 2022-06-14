package repository

import (
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Thought
}

type Authorization interface {
}

type Thought interface {
	Create(input entity.ThoughtInput) (entity.ThoughtCreateResponse, error)
	RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error)
	RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtPassphraseInput, error)
	CheckThoughtExists(thoughtKey string) (bool, error)
	BurnThought(thoughtKey, passphrase string) (bool, error)
	GetPassphraseOfThought(thoughtKey string) (string, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		Thought:       NewThoughtSql(db),
	}
}

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
	Metadata(metadataKey string) (entity.ThoughtMetadataResponse, error)
	CheckThoughtExists(thoughtKey string) (bool, error)
	ShowThought(thoughtKey, passphrase string) (entity.AccessThoughtResponse, error)
	BurnThought(thoughtKey, passphrase string) (bool, error)
	GetPassphraseOfThought(thoughtKey string) (string, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		Thought:       NewThoughtSql(db),
	}
}

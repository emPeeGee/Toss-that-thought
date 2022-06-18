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
	Create(input entity.ThoughtCreateInput) (entity.ThoughtCreateResponse, error)
	RetrieveMetadata(metadataKey string) (entity.ThoughtMetadataResponse, error)
	RetrieveThought(thoughtKey, passphrase string) (entity.ThoughtResponse, error)
	CheckThoughtExists(thoughtKey string) (entity.ThoughtValidityInformation, error)
	CheckMetadataExists(metadataKey string) (bool, error)
	BurnThought(metadataKey, passphrase string) (bool, error)
	MarkAsViewed(thoughtKey, passphrase string) error
	GetPassphraseOfThoughtByMetadataKey(metadataKey string) (string, error)
	GetPassphraseOfThoughtByThoughtKey(thoughtKey string) (string, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		Thought:       NewThoughtSql(db),
	}
}

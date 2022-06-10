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
	Test() (interface{}, error)
	Create(input entity.ThoughtInput) (entity.ThoughtResponse, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		Thought:       NewThoughtSql(db),
	}
}

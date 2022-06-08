package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	Thought
}

type Authorization interface {
}

type Thought interface {
	Test() (interface{}, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		Thought:       NewThoughtSql(db),
	}
}

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

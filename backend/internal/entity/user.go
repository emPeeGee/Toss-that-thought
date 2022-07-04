package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password string `json:"password" gorm:"notNull;size:256"`
	Username string `json:"username" gorm:"notNull;unique;size:64"`
	Name     string `json:"name" gorm:"notNull;size:256"`
	Thoughts []Thought
}

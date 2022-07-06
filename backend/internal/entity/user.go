package entity

import (
	"github.com/emPeeee/ttt/pkg/crypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `json:"password" gorm:"notNull;size:256"`
	Username string `json:"username" gorm:"notNull;unique;size:64"`
	Name     string `json:"name" gorm:"notNull;size:256"`
	Thoughts []Thought
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := crypt.HashPassphrase(password);
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return nil
}

func (u *User) CheckPassword(password string) error {
	return crypt.CheckPasswordHashes(password, u.Password)
}

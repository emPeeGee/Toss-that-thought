package auth

import (
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	createUser(user entity.User) error
	getUserById(id uint) (UserResponse, error)
	getUserByUsername(username string) (entity.User, error)
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewAuthRepository(db *gorm.DB, logger log.Logger) *repository {
	return &repository{db: db, logger: logger}
}

func (r *repository) createUser(user entity.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) getUserByUsername(username string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *repository) getUserById(id uint) (UserResponse, error) {
	var user UserResponse

	if err := r.db.Model(&entity.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return UserResponse{}, err
	}

	return user, nil
}

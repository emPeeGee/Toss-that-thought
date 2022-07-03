package auth

import (
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	createUser(user createUserDTO) error
	getUserByUsername(username string) (entity.User, error)
	getHashedPasswordByUsername(username string) (userHashedPassword, error)
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewAuthRepository(db *gorm.DB, logger log.Logger) *repository {
	return &repository{db: db, logger: logger}
}

func (r *repository) createUser(user createUserDTO) error {
	// maybe dto is not needed here?
	newUser := entity.User{
		Username: user.Username,
		Password: user.Password,
		Name:     user.Name,
	}

	if err := r.db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) getUserByUsername(username string) (entity.User, error) {
	var user entity.User
	// TODO if First if first when where does not have effect

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *repository) getHashedPasswordByUsername(username string) (userHashedPassword, error) {
	var userPassword userHashedPassword

	if err := r.db.Model(&entity.User{}).Where("username = ?", username).First(&userPassword).Error; err != nil {
		return userHashedPassword{}, err
	}

	return userPassword, nil
}

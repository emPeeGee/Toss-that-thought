package auth

import (
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user CreateUserDTO) error
	GetUserByUsername(username string) (entity.User, error)
	GetHashedPasswordByUsername(username string) (UserHashedPassword, error)
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewAuthRepository(db *gorm.DB, logger log.Logger) *repository {
	return &repository{db: db, logger: logger}
}

func (r *repository) CreateUser(user CreateUserDTO) error {
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

func (r *repository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	// TODO if First if first when where does not have effect

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *repository) GetHashedPasswordByUsername(username string) (UserHashedPassword, error) {
	var userHashedPassword UserHashedPassword

	if err := r.db.Model(&entity.User{}).Where("username = ?", username).First(&userHashedPassword).Error; err != nil {
		return UserHashedPassword{}, err
	}

	return userHashedPassword, nil
}

package auth

import (
	"encoding/json"
	"github.com/emPeeee/ttt/internal/config"
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Service interface {
	createUser(input createUserDTO) error
	generateToken(credentials credentialsDTO) (string, error)
	getUserById(id uint) (UserResponse, error)
}

type service struct {
	repo   Repository
	cfg    *config.Auth
	logger log.Logger
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"userId"`
}

func NewAuthService(repository Repository, cfg *config.Auth, logger log.Logger) *service {
	return &service{repo: repository, cfg: cfg, logger: logger}
}

func (s *service) createUser(userDTO createUserDTO) error {
	user := entity.User{
		Username: userDTO.Username,
		Name:     userDTO.Name,
	}

	err := user.HashPassword(userDTO.Password)
	if err != nil {
		return err
	}

	err = s.repo.createUser(user)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) generateToken(credentials credentialsDTO) (string, error) {
	user, err := s.repo.getUserByUsername(credentials.Username)
	if err != nil {
		return "such user does not exist", err
	}

	err = user.CheckPassword(credentials.Password)
	if err != nil {
		return "password does not match", err
	}

	str, _ := json.MarshalIndent(user, "", "\t")
	s.logger.Debug(string(str))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.cfg.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(s.cfg.SigningKey))
}

func (s *service) getUserById(id uint) (UserResponse, error) {
	return s.repo.getUserById(id)
}

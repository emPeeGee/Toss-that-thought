package auth

import "time"

type createUserDTO struct {
	Password string `json:"password" validate:"required,min=4,max=256"`
	Username string `json:"username" validate:"required,min=3,max=64"`
	Name     string `json:"name" validate:"required,min=2,max=64"`
}

type credentialsDTO struct {
	Password string `json:"password" validate:"required,min=4,max=256"`
	Username string `json:"username" validate:"required,min=3,max=64"`
}

type userHashedPassword struct {
	Password string
}

type UserResponse struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

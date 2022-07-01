package auth

type CreateUserDTO struct {
	Password string `json:"password" validate:"required,min=4,max=256"`
	Username string `json:"username" validate:"required,min=3,max=64"`
	Name     string `json:"name" validate:"required,min=2,max=64"`
}

type CredentialsDTO struct {
	Password string `json:"password" validate:"required,min=4,max=256"`
	Username string `json:"username" validate:"required,min=3,max=64"`
}

type UserHashedPassword struct {
	Password string
}
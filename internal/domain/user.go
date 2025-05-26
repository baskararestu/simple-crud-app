package domain

import "time"

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
LastAccessLogin time.Time `json:"last_access_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user User) error
	FindAll() ([]User, error)
	Update(id int, user User) error
	Delete (id int) error
	GetUserByEmail(email string) (User, error)
}


type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	RoleID string `json:"role_id" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserDataResponse struct{
	RoleID string `json:"role_id"`
	RoleName string `json:"role_name"`
	Name string `json:"name"`
	Email string `json:"email"`
	LastAccess string `json:"last_access"`
}

type UpdateUserRequest struct {
Name     string `json:"name" validate:"required"`
}
package model

import (
	"fmt"
	"time"

	"github.com/eduardonakaidev/task-flow/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// User representa um usuário do sistema.
type User struct {
	ID        uuid.UUID `json:"id"`                                           // Identificador único
	Username  string    `json:"username" validate:"required,max=80"`          // Nome de usuário
	Email     string    `json:"email" validate:"required,email"`              // Email válido
	Password  string    `json:"password" validate:"required,min=8,max=80"`    // Senha entre 8 e 80 caracteres
	CreatedAt time.Time `json:"created_at"`                                   // Data de criação
	UpdatedAt time.Time `json:"updated_at"`                                   // Data de atualização
}

// Validator instância o validador global.
var validate = validator.New()

func CreateUser(username, email, password string) (*User, error) {
		// Criando o objeto User
		user := &User{
			ID:        uuid.New(),
			Username:  username,
			Email:     email,
			Password:  password, // Temporariamente armazenada "pura"
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	
		// Validação dos campos usando o validator
		if err := validate.Struct(user); err != nil {
			return nil, err
		}
	
		// Hash da senha antes de retornar
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password")
		}
		user.Password = hashedPassword
	
		return user, nil
}

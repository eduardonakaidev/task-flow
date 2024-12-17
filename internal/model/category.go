package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `json:"id"`          // Identificador único do tipo UUID
	Name        string    `json:"name"`        // Nome da categoria
	Description string    `json:"description"` // Descrição opcional
	CreatedAt   time.Time `json:"created_at"`  // Data de criação
	UpdatedAt   time.Time `json:"updated_at"`  // Data de atualização
}

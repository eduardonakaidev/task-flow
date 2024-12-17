package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`          // Identificador único do tipo UUID
	Title       string    `json:"title"`       // Título da tarefa
	Description string    `json:"description"` // Descrição opcional
	DueDate     time.Time `json:"due_date"`    // Data de vencimento
	Completed   bool      `json:"completed"`   // Status de conclusão da tarefa
	CategoryID  uuid.UUID `json:"category_id"` // Chave estrangeira para a Category
	OwnerID     uuid.UUID `json:"owner_id"`      // Chave estrangeira para o User (dono da tarefa)
	CreatedAt   time.Time `json:"created_at"`  // Data de criação
	UpdatedAt   time.Time `json:"updated_at"`  // Data de atualização
}
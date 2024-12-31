package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
)

// Category representa uma categoria de tarefas.
type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// User representa um usuário do sistema.
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Task representa uma tarefa.
type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed"`
	CategoryID  uuid.UUID `gorm:"type:uuid;not null" json:"category_id"`
	OwnerID     uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relacionamentos
	Category Category `gorm:"foreignKey:CategoryID"`
	Owner    User     `gorm:"foreignKey:OwnerID"`
}

// ConnectDatabase conecta ao banco e executa as migrações.
func ConnectDatabase() *gorm.DB {
	// Conectar ao SQLite
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrar os modelos
	err = db.AutoMigrate(&Category{}, &User{}, &Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully!")
	return db
}

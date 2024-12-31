package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/eduardonakaidev/task-flow/internal/config"
	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("POST /api/task", func(w http.ResponseWriter, r *http.Request) {
		db := config.ConnectDatabase()

		// Exemplo: Criar uma nova categoria
		newCategory := config.Category{
			ID:          uuid.New(),
			Name:        "Work",
			Description: "Tasks related to work",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		if err := db.Create(&newCategory).Error; err != nil {
			http.Error(w, "Failed to create category", http.StatusInternalServerError)
			log.Println("Failed to create category:", err)
			return
		}
		// Responder com sucesso
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		
		response := map[string]interface{}{
			"message": "Category created successfully",
			"data":    newCategory,
		}
		json.NewEncoder(w).Encode(response)
		log.Println("Category created successfully!")
	})
	http.ListenAndServe(":3000", nil)
}

package handlers

import (
	"fmt"
	"net/http"

	"example.com/internal/database"
)

type Handlers struct {
	db *database.Database
}

func NewHandlers(db *database.Database) *Handlers {
	return &Handlers{db: db}
}

func (h *Handlers) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		test := h.db.DBRepo.GetAll()
		fmt.Println(test)
	}
}

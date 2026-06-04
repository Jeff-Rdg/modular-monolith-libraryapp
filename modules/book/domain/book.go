package domain

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Available bool      `json:"available"`
	CreatedAt time.Time `json:"created_at"`
}

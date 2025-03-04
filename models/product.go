package models

import (
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	Name       	string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}

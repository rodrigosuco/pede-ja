package models

import "time"

type Product struct {
	ID          string    `json:"id"`
	Name       	string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}

var Products = []Product{
	{ID: "1", Name: "Product 1", Description: "Product 1", CreatedAt: time.Now(), Status: "deactivated"},
	{ID: "2", Name: "Product 2", Description: "Product 2", CreatedAt: time.Now().AddDate(0, 0, 1), Status: "active"},
	{ID: "3", Name: "Product 3", Description: "Product 3", CreatedAt: time.Now().AddDate(0, 0, 2), Status: "active"},
}
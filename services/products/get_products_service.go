package services

import (
	"context"

	"github.com/rodrigosuco/pede-ja/models"
	"github.com/rodrigosuco/pede-ja/internal/database"
)

func GetAllProducts() ([]models.Product, error) {
	query := `SELECT id, name, description, created_at, status FROM products;`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.Status)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

package services

import (
	"context"
	"fmt"

	"github.com/rodrigosuco/pede-ja/internal/database"
	"github.com/rodrigosuco/pede-ja/models"
)

func CreateProduct(product models.Product) (models.Product, error)  {
	query := `INSERT INTO products (name, description, status)
						VALUES ($1, $2, $3)
						RETURNING id, name, description, created_at, status;`

	err := database.DB.QueryRow(
		context.Background(),
		query,
		product.Name,
		product.Description,
		product.Status,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.CreatedAt,
		&product.Status,
	)

	if err != nil {
		return models.Product{}, fmt.Errorf("failed to create product: %v", err)
	}

	return product, nil

}

package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategory(
	db *sql.DB,
) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name, description string) (*Category, error) {

	cat := Category{
		Name:        name,
		Description: description,
		ID:          uuid.NewString(),
	}

	_, err := c.db.Exec(
		"INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		cat.ID,
		cat.Name,
		cat.Description,
	)

	if !errors.Is(err, nil) {
		return nil, err
	}

	return &cat, nil
}

func (c *Category) FindAll() (*[]Category, error) {

	rows, err := c.db.Query("SELECT id, name, description FROM categories")

	if !errors.Is(err, nil) {
		return nil, err
	}

	var categories []Category

	for rows.Next() {

		var cat Category
		rows.Scan(&cat.ID, &cat.Name, &cat.Description)

		categories = append(categories, cat)
	}

	return &categories, nil
}

func (c *Category) FindById(id string) (*Category, error) {

	row, err := c.db.Query("SELECT * FROM categories WHERE id = $1 LIMIT 1", id)

	if !errors.Is(err, nil) {
		return nil, err
	}

	var cat Category

	fmt.Println(row.Columns())

	for row.Next() {
		row.Scan(&cat.ID, &cat.Name, &cat.Description)
	}

	return &cat, nil
}

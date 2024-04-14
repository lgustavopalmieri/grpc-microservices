package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) CreateCategory(name string, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec(
		"INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description,
	)
	if err != nil {
		return Category{}, err
	}
	return Category{
			ID:          id,
			Name:        name,
			Description: description,
		},
		nil
}

func (c *Category) Find(id string) (*Category, error) {
	query, err := c.db.Query("SELECT id, name, description FROM categories WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var category Category
	if query.Next() {
		err := query.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		return &category, nil
	}

	return nil, errors.New("category not found")
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) UpdateCategory(categoryID string, newName string, newDescription string) (Category, error) {
	stmt, err := c.db.Prepare("UPDATE categories SET name = ?, description = ? WHERE id = ?")
	if err != nil {
		return Category{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newName, newDescription, categoryID)
	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          categoryID,
		Name:        newName,
		Description: newDescription,
	}, nil
}


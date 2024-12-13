package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gpereverzev/CashWise/src/data"
	"strings"
)

type categoriesDB struct {
	db *sql.DB
}

func newCategoriesDB(db *sql.DB) *categoriesDB {
	return &categoriesDB{db: db}
}

func (c *categoriesDB) Get(categoryID int) (*data.Category, error) {
	category := &data.Category{}

	query := `SELECT categoryID, userID, name, description  
			FROM Categories 
			WHERE categoryID = $1`
	err := c.db.QueryRow(query, categoryID).Scan(&category.CategoryID, &category.UserID, &category.Name,
		&category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category not found")
		}
		return nil, err
	}

	return category, nil
}

func (c *categoriesDB) Insert(category *data.Category) error {
	query := `
		INSERT INTO Categories (categoryID, userID, name, description)
		VALUES ($1, $2, $3, $4)
	`
	_, err := c.db.Exec(query, category.CategoryID, category.UserID, category.Name, category.Description)
	return err
}

func (c *categoriesDB) Update(categoryID int, updateFields map[string]interface{}) error {
	var setClauses []string
	var args []interface{}
	i := 1

	for field, value := range updateFields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}
	args = append(args, categoryID)

	query := fmt.Sprintf("UPDATE Categories SET %s WHERE categoryID = $%d",
		strings.Join(setClauses, ", "), i)
	_, err := c.db.Exec(query, args...)
	return err
}

func (c *categoriesDB) Delete(categoryID int) error {
	query := `DELETE FROM Categories WHERE categoryID = $1`
	_, err := c.db.Exec(query, categoryID)
	return err
}

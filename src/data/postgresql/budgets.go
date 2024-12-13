package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gpereverzev/CashWise/src/data"
	"strings"
)

type budgetsDB struct {
	db *sql.DB
}

func newBudgetsDB(db *sql.DB) *budgetsDB {
	return &budgetsDB{db: db}
}

func (b *budgetsDB) Get(budgetID int) (*data.Budget, error) {
	budget := &data.Budget{}

	query := `SELECT budgetID, userID, title, initialBalance, limit, period  
			FROM Budgets 
			WHERE budgetID = $1`
	err := b.db.QueryRow(query, budgetID).Scan(&budget.BudgetID, &budget.UserID, &budget.Title,
		&budget.InitialBalance, &budget.Limit, &budget.Period)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("budget not found")
		}
		return nil, err
	}

	return budget, nil
}

func (b *budgetsDB) Insert(budget *data.Budget) error {
	query := `
		INSERT INTO Budgets (budgetID, userID, title, initialBalance, limit, period)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := b.db.Exec(query, budget.BudgetID, budget.UserID, budget.Title, budget.InitialBalance, budget.Limit, budget.Period)
	return err
}

func (b *budgetsDB) Update(budgetID int, updateFields map[string]interface{}) error {
	var setClauses []string
	var args []interface{}
	i := 1

	for field, value := range updateFields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}
	args = append(args, budgetID)

	query := fmt.Sprintf("UPDATE Budgets SET %s WHERE budgetID = $%d",
		strings.Join(setClauses, ", "), i)
	_, err := b.db.Exec(query, args...)
	return err
}

func (b *budgetsDB) Delete(budgetID int) error {
	query := `DELETE FROM Budgets WHERE budgetID = $1`
	_, err := b.db.Exec(query, budgetID)
	return err
}

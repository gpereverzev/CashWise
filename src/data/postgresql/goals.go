package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gpereverzev/CashWise/src/data"
	"strings"
)

type goalsDB struct {
	db *sql.DB
}

func newGoalsDB(db *sql.DB) *goalsDB {
	return &goalsDB{db: db}
}

func (g *goalsDB) Get(goalID int) (*data.Goal, error) {
	goal := &data.Goal{}

	query := `SELECT goalID, userID, title, targetAmount, currentAmount, deadline, status  
			FROM Goals 
			WHERE goalID = $1`
	err := g.db.QueryRow(query, goalID).Scan(&goal.GoalID, &goal.UserID, &goal.Title,
		&goal.TargetAmount, &goal.CurrentAmount, &goal.Deadline, &goal.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("goal not found")
		}
		return nil, err
	}

	return goal, nil
}

func (g *goalsDB) Insert(goal *data.Goal) error {
	query := `
		INSERT INTO Goals (goalID, userID, title, targetAmount, currentAmount, deadline, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := g.db.Exec(query, goal.GoalID, goal.UserID, goal.Title, goal.TargetAmount, goal.CurrentAmount, goal.Deadline, goal.Status)
	return err
}

func (g *goalsDB) Update(goalID int, updateFields map[string]interface{}) error {
	var setClauses []string
	var args []interface{}
	i := 1

	for field, value := range updateFields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}
	args = append(args, goalID)

	query := fmt.Sprintf("UPDATE Goals SET %s WHERE goalID = $%d",
		strings.Join(setClauses, ", "), i)
	_, err := g.db.Exec(query, args...)
	return err
}

func (g *goalsDB) Delete(goalID int) error {
	query := `DELETE FROM Goals WHERE goalID = $1`
	_, err := g.db.Exec(query, goalID)
	return err
}

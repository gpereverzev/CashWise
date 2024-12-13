package data

import "time"

type GoalsDB interface {
	Get(id int) (*Goal, error)
	Insert(user *Goal) error
	Update(id int, updateFields map[string]interface{}) error
	Delete(id int) error
}

type Goal struct {
	GoalID        int       `json:"_id"`
	UserID        int       `json:"user_id"`
	Title         string    `json:"title"`
	TargetAmount  float32   `json:"target_amount"`
	CurrentAmount float32   `json:"current_amount"`
	Deadline      time.Time `json:"deadline"`
	Status        string    `json:"status"`
}

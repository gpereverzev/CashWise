package data

type BudgetsDB interface {
	Get(id int) (*Budget, error)
	Insert(user *Budget) error
	Update(id int, updateFields map[string]interface{}) error
	Delete(id int) error
}

type Budget struct {
	BudgetID       int     `json:"_id"`
	UserID         int     `json:"user_id"`
	Title          string  `json:"title"`
	InitialBalance float32 `json:"initial_balance"`
	Limit          float32 `json:"limit"`
	Period         string  `json:"period"`
}

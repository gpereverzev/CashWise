package data

type TransactionsDB interface {
	Get(id int) (*Transaction, error)
	Insert(user *Transaction) error
	Update(id int, updateFields map[string]interface{}) error
	Delete(id int) error
}

type Transaction struct {
	TransactionID int     `json:"_id"`
	UserID        int     `json:"user_id"`
	CategoryID    string  `json:"category_id"`
	Type          string  `json:"type"`
	Amount        float32 `json:"amount"`
	Date          string  `json:"date"`
	Description   string  `json:"description"`
}

package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gpereverzev/CashWise/src/data"
	"strings"
)

type transactionsDB struct {
	db *sql.DB
}

func newTransactionsDB(db *sql.DB) *transactionsDB {
	return &transactionsDB{db: db}
}

func (t *transactionsDB) Get(transactionID int) (*data.Transaction, error) {
	transaction := &data.Transaction{}

	query := `SELECT transactionID, userID, categoryID, type, amount, date_, description  
				FROM Transactions 
	          WHERE transactionID = $1`
	err := t.db.QueryRow(query, transactionID).Scan(&transaction.TransactionID, &transaction.UserID, &transaction.CategoryID,
		&transaction.Type, &transaction.Amount, &transaction.Date, &transaction.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return transaction, nil
}

func (t *transactionsDB) Insert(transaction *data.Transaction) error {
	query := `
		INSERT INTO Transactions (transactionID, userID, categoryID, type, amount, date_, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := t.db.Exec(query, transaction.TransactionID, &transaction.UserID, &transaction.CategoryID, &transaction.Type, &transaction.Amount, &transaction.Date, &transaction.Description)
	return err
}

func (t *transactionsDB) Update(transactionID int, updateFields map[string]interface{}) error {
	var setClauses []string
	var args []interface{}
	i := 1

	for field, value := range updateFields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}
	args = append(args, transactionID)

	query := fmt.Sprintf("UPDATE Transactions SET %s WHERE transactionID = $%d",
		strings.Join(setClauses, ", "), i)
	_, err := t.db.Exec(query, args...)
	return err
}

func (t *transactionsDB) Delete(transactionID int) error {
	query := `DELETE FROM Transactions WHERE transactionID = $1`
	_, err := t.db.Exec(query, transactionID)
	return err
}

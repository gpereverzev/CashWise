package postgresql

import (
	"database/sql"
	"github.com/gpereverzev/CashWise/src/data"
)

type masterDB struct {
	users        *usersDB
	transactions *transactionsDB
	budgets      *budgetsDB
	categories   *categoriesDB
	goals        *goalsDB
}

func NewMasterDB(db *sql.DB) data.MasterDB {
	return &masterDB{
		users:        newUsersDB(db),
		transactions: newTransactionsDB(db),
		budgets:      newBudgetsDB(db),
		categories:   newCategoriesDB(db),
		goals:        newGoalsDB(db),
	}
}

func (m *masterDB) Users() data.UsersDB {
	return m.users
}

func (m *masterDB) Transactions() data.TransactionsDB {
	return m.transactions
}

func (m *masterDB) Budgets() data.BudgetsDB {
	return m.budgets
}

func (m *masterDB) Categories() data.CategoriesDB {
	return m.categories
}

func (m *masterDB) Goals() data.GoalsDB {
	return m.goals
}

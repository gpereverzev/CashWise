package data

type MasterDB interface {
	Users() UsersDB
	Transactions() TransactionsDB
	Budgets() BudgetsDB
	Categories() CategoriesDB
}

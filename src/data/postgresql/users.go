package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gpereverzev/CashWise/src/data"
	"strings"
)

type usersDB struct {
	db *sql.DB
}

func newUsersDB(db *sql.DB) *usersDB {
	return &usersDB{db: db}
}

func (u *usersDB) Get(userID int) (*data.User, error) {
	user := &data.User{}

	query := `SELECT userID, firstName, lastName, email, passwordHash, role_
	          FROM "Users" 
	          WHERE userID = $1`
	err := u.db.QueryRow(query, userID).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (u *usersDB) Insert(user *data.User) error {
	query := `
		INSERT INTO "Users" ( firstName, lastName, email, passwordHash, role_)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := u.db.Exec(query, user.FirstName, user.LastName, user.Email, user.PasswordHash, user.Role)
	return err
}

func (u *usersDB) Update(userID int, updateFields map[string]interface{}) error {
	var setClauses []string
	var args []interface{}
	i := 1

	for field, value := range updateFields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}
	args = append(args, userID)

	query := fmt.Sprintf("UPDATE \"Users\" SET %s WHERE userID = $%d",
		strings.Join(setClauses, ", "), i)
	_, err := u.db.Exec(query, args...)
	return err
}

func (u *usersDB) FindByEmail(email string) (*data.User, error) {
	user := &data.User{}

	query := `SELECT userID, firstName, lastName, email, passwordHash, role_
    		FROM "Users" 
    		WHERE email = $1`
	err := u.db.QueryRow(query, email).Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (u *usersDB) Delete(userID int) error {
	query := `DELETE FROM "Users" WHERE userID = $1`
	_, err := u.db.Exec(query, userID)
	return err
}

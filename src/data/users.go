package data

type UsersDB interface {
	Get(id int) (*User, error)
	Insert(user *User) error
	Update(id int, updateFields map[string]interface{}) error
	Delete(id int) error
	FindByEmail(email string) (*User, error)
}

type User struct {
	UserID       int    `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

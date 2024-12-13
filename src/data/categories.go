package data

type CategoriesDB interface {
	Get(id int) (*Category, error)
	Insert(user *Category) error
	Update(id int, updateFields map[string]interface{}) error
	Delete(id int) error
}

type Category struct {
	CategoryID  int    `json:"_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

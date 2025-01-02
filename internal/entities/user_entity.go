package entities

type User struct {
	ID      int
	Name    string
	Balance float64
}

type IUserRepository interface {
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) error
}

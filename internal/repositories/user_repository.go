package repositories

import (
	"fmt"
	"paper-disbursement/internal/entities"
)

type userRepository struct {
	users map[int]*entities.User
}

func NewUserRepository() entities.IUserRepository {
	return &userRepository{
		users: map[int]*entities.User{
			1: {ID: 1, Name: "Alice", Balance: 100.0},
			2: {ID: 2, Name: "Bob", Balance: 200.0},
		},
	}
}

func (r *userRepository) GetUserByID(id int) (*entities.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *entities.User) error {
	if _, exists := r.users[user.ID]; !exists {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}

package repositories_test

import (
	"paper-disbursement/internal/entities"
	"paper-disbursement/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_GetUserByID(t *testing.T) {
	repo := repositories.NewUserRepository()

	t.Run("User exists", func(t *testing.T) {
		user, err := repo.GetUserByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "Alice", user.Name)
	})

	t.Run("User does not exist", func(t *testing.T) {
		user, err := repo.GetUserByID(99)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
	})
}

func TestRepository_UpdateUser(t *testing.T) {
	repo := repositories.NewUserRepository()

	t.Run("Update existing user", func(t *testing.T) {
		user := &entities.User{ID: 1, Name: "Alice", Balance: 50.0}
		err := repo.UpdateUser(user)
		assert.NoError(t, err)

		updatedUser, _ := repo.GetUserByID(1)
		assert.Equal(t, 50.0, updatedUser.Balance)
	})

	t.Run("Update non-existing user", func(t *testing.T) {
		user := &entities.User{ID: 99, Name: "Ghost", Balance: 100.0}
		err := repo.UpdateUser(user)
		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")
	})
}

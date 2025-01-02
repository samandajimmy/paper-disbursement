package usecases_test

import (
	"paper-disbursement/internal/repositories"
	"paper-disbursement/internal/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsecase_Disburse(t *testing.T) {
	repo := repositories.NewUserRepository()
	uc := usecases.NewDisburseUsecase(repo)

	t.Run("Successful disbursement", func(t *testing.T) {
		resp, err := uc.Disburse(1, 50.0)
		assert.NoError(t, err)
		assert.Equal(t, "Disbursement of 50.00 successful", resp.Message)
		assert.Equal(t, 50.0, resp.Balance)
	})

	t.Run("User not found", func(t *testing.T) {
		_, err := uc.Disburse(99, 50.0)
		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")
	})

	t.Run("Insufficient balance", func(t *testing.T) {
		_, err := uc.Disburse(1, 200.0)
		assert.Error(t, err)
		assert.EqualError(t, err, "insufficient balance")
	})

	t.Run("Invalid amount", func(t *testing.T) {
		_, err := uc.Disburse(1, -10.0)
		assert.Error(t, err)
		assert.EqualError(t, err, "amount must be greater than zero")
	})
}

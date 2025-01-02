package usecases

import (
	"fmt"
	"paper-disbursement/internal/entities"
)

type DisburseUsecase struct {
	IUserRepo entities.IUserRepository
}

func NewDisburseUsecase(iUserRepo entities.IUserRepository) entities.IDisburseUsecase {
	return &DisburseUsecase{IUserRepo: iUserRepo}
}

func (u *DisburseUsecase) Disburse(userID int, amount float64) (entities.DisbursementResponse, error) {
	if amount <= 0 {
		return entities.DisbursementResponse{}, fmt.Errorf("amount must be greater than zero")
	}

	user, err := u.IUserRepo.GetUserByID(userID)
	if err != nil {
		return entities.DisbursementResponse{}, err
	}

	if user.Balance < amount {
		return entities.DisbursementResponse{}, fmt.Errorf("insufficient balance")
	}

	user.Balance -= amount
	_ = u.IUserRepo.UpdateUser(user)

	return entities.DisbursementResponse{
		Message: fmt.Sprintf("Disbursement of %.2f successful", amount),
		Balance: user.Balance,
	}, nil
}

package entities

type DisbursementRequest struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

type DisbursementResponse struct {
	Message string  `json:"message"`
	Balance float64 `json:"balance"`
}

type IDisburseUsecase interface {
	Disburse(userID int, amount float64) (DisbursementResponse, error)
}

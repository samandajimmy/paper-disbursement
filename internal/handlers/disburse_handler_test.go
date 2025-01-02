package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"paper-disbursement/internal/entities"
	"paper-disbursement/internal/handlers"
	"paper-disbursement/internal/repositories"
	"paper-disbursement/internal/usecases"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDisburseHandler(t *testing.T) {
	repo := repositories.NewUserRepository()
	uc := usecases.NewDisburseUsecase(repo)
	handler := handlers.NewDisburseHandler(uc)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/disburse", handler.DisburseHandler)

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful disbursement",
			requestBody:    `{"user_id": 1, "amount": 50}`,
			expectedStatus: http.StatusOK,
			expectedBody:   "Disbursement of 50.00 successful",
		},
		{
			name:           "User not found",
			requestBody:    `{"user_id": 99, "amount": 50}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "user not found",
		},
		{
			name:           "Insufficient balance",
			requestBody:    `{"user_id": 1, "amount": 150}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "insufficient balance",
		},
		{
			name:           "Invalid amount",
			requestBody:    `{"user_id": 1, "amount": -10}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "amount must be greater than zero",
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{"user_id": 1, "amount": "abc"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid request payload",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/disburse", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			resp := w.Result()
			fmt.Println(resp)
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if resp.StatusCode == http.StatusOK {
				var response entities.DisbursementResponse
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Contains(t, response.Message, tt.expectedBody)
			} else {
				var errMessage gin.H
				err := json.NewDecoder(resp.Body).Decode(&errMessage)
				assert.NoError(t, err)
				assert.Contains(t, errMessage["error"], tt.expectedBody)
			}
		})
	}
}

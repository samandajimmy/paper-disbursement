package handlers

import (
	"net/http"
	"paper-disbursement/internal/entities"

	"github.com/gin-gonic/gin"
)

type DisburseHandler struct {
	IDisburseUsecase entities.IDisburseUsecase
}

func NewDisburseHandler(iDisburseUsecase entities.IDisburseUsecase) *DisburseHandler {
	return &DisburseHandler{IDisburseUsecase: iDisburseUsecase}
}

func (h *DisburseHandler) DisburseHandler(c *gin.Context) {
	var req entities.DisbursementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	resp, err := h.IDisburseUsecase.Disburse(req.UserID, req.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

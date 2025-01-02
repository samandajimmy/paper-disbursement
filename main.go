package main

import (
	"fmt"
	"paper-disbursement/internal/handlers"
	"paper-disbursement/internal/repositories"
	"paper-disbursement/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepo := repositories.NewUserRepository()
	disburseUsecase := usecases.NewDisburseUsecase(userRepo)
	disburseHandler := handlers.NewDisburseHandler(disburseUsecase)

	r := gin.Default()
	r.POST("/disburse", disburseHandler.DisburseHandler)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}

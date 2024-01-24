package usecases

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goEvent/entities"
)

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s succesfull", op),
		"data":    data,
	})
}

type CreateAppointmentResponse struct {
	Message string                       `json:"message"`
	Data    entities.AppointmentResponse `json:"data "`
}

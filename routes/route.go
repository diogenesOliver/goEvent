package routes

import (
	"github.com/gin-gonic/gin"
	usecases "github.com/goEvent/useCases"
)

func AppRoutes(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/v1")
	{
		v1.POST("/create/appointment", usecases.CreateAppointments)
		v1.GET("/appointments", usecases.GetAllAppointments)
	}
	return v1
}

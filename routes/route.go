package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/goEvent/useCases"
	usecases "github.com/goEvent/useCases"
)

func AppRoutes(r *gin.Engine) *gin.RouterGroup {
	appointmentController := controllers.NewAppointment()

	v1 := r.Group("/v1")
	{
		v1.POST("/events", appointmentController.CreateAppointments)
		v1.GET("/appointments", usecases.GetAllAppointments)
	}
	return v1
}

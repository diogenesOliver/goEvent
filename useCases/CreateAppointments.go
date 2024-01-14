package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goEvent/entities"
)

func (a *AppointmentController) CreateAppointments(ctx *gin.Context) {
	appointment := entities.Appointment{}
	if err := ctx.BindJSON(&appointment); err != nil {
		return
	}

	a.appointments = append(a.appointments, appointment)
	ctx.JSON(http.StatusOK, a.appointments)
}

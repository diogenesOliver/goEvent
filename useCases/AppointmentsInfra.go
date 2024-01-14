package usecases

import "github.com/goEvent/entities"

type AppointmentController struct {
	appointments []entities.Appointment
}

func NewAppointment() *AppointmentController {
	return &AppointmentController{}
}

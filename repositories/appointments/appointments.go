package appointments

import (
	"context"

	"github.com/goEvent/entities"
)

type AppointmentsRepository interface {
	Create(ctx context.Context, newAppointment entities.Appointment) error
	GetAll(ctx context.Context) ([]entities.Appointment, error)
}

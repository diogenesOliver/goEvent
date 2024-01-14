package entities

import (
	"github.com/pborman/uuid"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	ID string `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Interval    int    `json:"interval"`
}

func NewAppointment() *Appointment {
	appointment := Appointment{
		ID: uuid.New(),
	}
	return &appointment
}

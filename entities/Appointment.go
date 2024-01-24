package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Appointment struct {
	gorm.Model
	Title       string
	Description string
	Date        string
	Interval    int
}

type AppointmentResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
	Interval    int       `json:"interval"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

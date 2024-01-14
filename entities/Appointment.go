package entities

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Interval    int    `json:"interval"`
}

package entities

import "time"

type Appointment struct {
	ID          int        `gorm:"primary_key" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
	Interval    int        `json:"interval"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdateAt    *time.Time `db:"update_at" json:"update_at"`
}

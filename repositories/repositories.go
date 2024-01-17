package repositories

import (
	"github.com/goEvent/repositories/appointments"
	"github.com/jinzhu/gorm"
)

type Container struct {
	Appointment appointments.AppointmentsRepository
}

type Options struct {
	WriterGorm *gorm.DB
	ReaderGorm *gorm.DB
}

func New(opts Options) *Container {
	return &Container{
		Appointment: appointments.NewGormRepository(opts.WriterGorm, opts.ReaderGorm),
	}
}

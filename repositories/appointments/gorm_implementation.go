package appointments

import (
	"context"
	"errors"
	"fmt"

	"github.com/goEvent/entities"
	"github.com/jinzhu/gorm"
)

type repoGorm struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormRepository(writer, reader *gorm.DB) AppointmentsRepository {
	return &repoGorm{writer: writer, reader: reader}
}

func (repo *repoGorm) Create(ctx context.Context, newAppointment entities.Appointment) error {
	repo.writer.Table("appointments").Create(&newAppointment)

	if repo.writer.Error != nil {
		fmt.Println(repo.writer.Error)
		return errors.New("Appointment not registered")
	}
	return nil
}

func (repo *repoGorm) GetAll(ctx context.Context) ([]entities.Appointment, error) {
	var appointment []entities.Appointment
	repo.reader.Table("appointments").Find(&appointment)

	if repo.reader.Error != nil {
		fmt.Println(repo.reader.Error)
		return nil, errors.New("Appointment not found")
	}
	return appointment, nil
}

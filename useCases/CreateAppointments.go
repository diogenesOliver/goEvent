package usecases

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goEvent/configs"
	"github.com/goEvent/entities"
	"github.com/goEvent/repositories"
)

func CreateAppointments(ctx *gin.Context) {
	request := CreateAppointmentRequest{}
	ctx.BindJSON(&request)

	appointment := entities.Appointment{
		Title:       request.Title,
		Description: request.Description,
		Interval:    request.Interval,
		Date:        request.Date,
	}

	repos := repositories.New(repositories.Options{
		ReaderGorm: configs.GetReaderGorm(),
		WriterGorm: configs.GetWriterGorm(),
	})

	fmt.Println(appointment)

	err := repos.Appointment.Create(context.Background(), appointment)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, appointment)
}

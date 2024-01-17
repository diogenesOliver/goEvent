package usecases

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goEvent/configs"
	"github.com/goEvent/repositories"
)

func GetAllAppointments(ctx *gin.Context) {
	repos := repositories.New(repositories.Options{
		ReaderGorm: configs.GetReaderGorm(),
		WriterGorm: configs.GetWriterGorm(),
	})

	appointments, err := repos.Appointment.GetAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, appointments)
}

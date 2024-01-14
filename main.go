package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goEvent/routes"
)

func main() {
	app := gin.Default()
	routes.AppRoutes(app)

	err := app.Run(":8080")
	if err != nil {
		return
	}
}

package routes

import "github.com/gin-gonic/gin"

func AppRoutes(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/v1")
	{
		v1.GET("/events", func(ctx *gin.Context) {
			ctx.JSON(200, "This is a simple event")
		})
	}
	return v1
}

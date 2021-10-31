package root

import (
	"github.com/gin-gonic/gin"
	"petshop/handler"
)

func SetupRouting() *gin.Engine {
	route := gin.New()

	userRoute := route.Group("/user")
	{
		userRoute.GET("/get/:id", handler.GetUser())
		userRoute.POST("/list", handler.GetUsers())
	}

	return route
}

package Routes

import (
	"crudTask/Controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	route := gin.Default()
	group1 := route.Group("contact-api")
	{
		group1.POST("contact", Controllers.CreateConact)
	}
	return route
}

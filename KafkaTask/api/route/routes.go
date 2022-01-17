package route

import (
	"KafkaTask/api/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	route := gin.Default()
	group1 := route.Group("contact-api")
	{
		group1.POST("contact", controller.PostContact)
		group1.GET("contact", controller.GetContacts)
	}
	return route
}

package route

import (
	"kafkamicroservice/consumer/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	route := gin.Default()
	group1 := route.Group("contact-api")
	{
		group1.GET("contact/:id", controller.GetContact)
		group1.GET("contact/", controller.ListContacts)

	}
	return route
}

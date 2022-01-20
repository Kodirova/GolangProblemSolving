package routes

import (
	"kafkamicroservice/api/controller"
	// _ "kafkamicroservice/api"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	route := gin.Default()
	group1 := route.Group("contact-api")
	{
		group1.POST("contact", controller.PostContact)
		group1.PUT("contact", controller.UpdateContact)
		group1.DELETE("contact", controller.DeleteContact)
		// group1.GET("contact", controller.GetContacts)
	}
	return route
}

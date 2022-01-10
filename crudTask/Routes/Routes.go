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
		group1.GET("contact", Controllers.ListContacts)
		group1.GET("contact/:id", Controllers.GetContact)
		group1.PUT("contact/:id", Controllers.UpdateContact)
		group1.DELETE("contact/:id", Controllers.DeleteContact)
	}
	return route
}

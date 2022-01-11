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
	group2 := route.Group("task-api")
	{
		group2.POST("task", Controllers.CreateTask)
		group2.GET("task", Controllers.ListTask)
		group2.GET("task/:id", Controllers.GetTask)
		group2.PUT("task/:id", Controllers.UpdateTask)
		group2.DELETE("task/:id", Controllers.DeleteTask)
	}
	return route
}

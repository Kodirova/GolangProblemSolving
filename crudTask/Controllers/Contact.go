package Controllers

import (
	"crudTask/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateConact(n *gin.Context) {
	var contact Models.Contact
	n.BindJSON(&contact)
	err := Models.CreateContact(&contact)
	if err != nil {
		fmt.Println(err.Error())
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, contact)
	}
}

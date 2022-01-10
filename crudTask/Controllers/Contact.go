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

func ListContacts(n *gin.Context) {
	var contact []Models.Contact
	err := Models.ListContacts(&contact)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, contact)
	}
}

func GetContact(n *gin.Context) {
	id := n.Params.ByName("id")
	var contact Models.Contact
	err := Models.GetContact(&contact, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, contact)
	}
}

func UpdateContact(n *gin.Context) {
	id := n.Params.ByName("id")
	var contact Models.Contact
	err := Models.GetContact(&contact, id)
	if err != nil {
		n.JSON(http.StatusNotFound, contact)
	}
	n.BindJSON(&contact)
	err = Models.UpdateContact(&contact, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, contact)
	}
}

func DeleteContact(n *gin.Context) {
	id := n.Params.ByName("id")
	var contact Models.Contact

	err := Models.DeleteContact(&contact, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, id)
	}
}

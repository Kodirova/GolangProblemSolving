package controller

import (
	"kafkamicroservice/consumer/models"
	"kafkamicroservice/consumer/query"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetContact(n *gin.Context) {
	id := n.Params.ByName("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	var contact models.Contact
	err1 := query.GetContact(&contact, id_int)
	if err1 != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, contact)
	}
}

func ListContacts(n *gin.Context) {
	var contact []models.Contact
	err := query.ListContacts(&contact)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, contact)
	}
}

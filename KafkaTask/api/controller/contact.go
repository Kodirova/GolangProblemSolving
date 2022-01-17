package controller

import (
	"KafkaTask/api/model"
	"KafkaTask/producer"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var contacts []model.Contact

func PostContact(c *gin.Context) {
	newContact := new(model.Contact)
	if err := c.BindJSON(&newContact); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	dataInBytes, err := json.Marshal(newContact)
	producer.PushHandlerToQueue("create", dataInBytes)
	if err != nil {
		response := gin.H{
			"status":  http.StatusNotFound,
			"message": "Message has not been sent.",
			"data":    newContact,
		}
		c.IndentedJSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  http.StatusOK,
			"message": "Message has been sent.",
			"data":    newContact,
		}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func GetContact(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

func GetContacts(c *gin.Context) {
	id := c.Param("id")
	for _, contact := range contacts {
		if contact.ID == id {
			c.IndentedJSON(http.StatusOK, contact)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

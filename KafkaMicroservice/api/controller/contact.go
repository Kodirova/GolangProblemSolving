package controller

import (
	"encoding/json"
	"fmt"
	"kafkamicroservice/api/models"
	"kafkamicroservice/api/producer"
	"kafkamicroservice/api/proxy"
	"log"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

var contacts []models.Contact

func PostContact(c *gin.Context) {
	newContact := new(models.Contact)
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
	var contact models.Contact
	proxy.MakeProxy(c, "http://localhost:8081/", "contact-api/contact/")
	c.IndentedJSON(http.StatusOK, contact)
}

// func GetContacts(c *gin.Context) {
// 	id := c.GetInt("id")
// 	for _, contact := range contacts {
// 		if contact.ID == id {
// 			c.IndentedJSON(http.StatusOK, contact)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

// }

func UpdateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.BindJSON(&contact); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	dataInBytes, err := json.Marshal(contact)
	log.Println(contact)
	if err != nil {
		response := gin.H{
			"status":  http.StatusNotFound,
			"message": "Message has not been sent.",
			"data":    contact,
		}
		c.IndentedJSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  http.StatusOK,
			"message": "Message has been sent.",
			"data":    contact,
		}
		c.IndentedJSON(http.StatusOK, response)
	}
	producer.PushHandlerToQueue("update", dataInBytes)

}

func DeleteContact(c *gin.Context) {
	var contact models.Contact
	if err := c.BindJSON(&contact); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	dataInBytes, err := json.Marshal(contact)
	log.Println(contact)
	if err != nil {
		response := gin.H{
			"status":  http.StatusNotFound,
			"message": "Message has not been sent.",
			"data":    contact,
		}
		c.IndentedJSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  http.StatusOK,
			"message": "Message has been sent.",
			"data":    contact,
		}
		c.IndentedJSON(http.StatusOK, response)
	}
	producer.PushHandlerToQueue("delete", dataInBytes)
}

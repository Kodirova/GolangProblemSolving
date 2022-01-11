package Controllers

import (
	"crudTask/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(n *gin.Context) {
	var task Models.Task
	n.BindJSON(&task)
	err := Models.CreateTask(&task)
	if err != nil {
		fmt.Println(err.Error())
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, task)
	}
}

func ListTask(n *gin.Context) {
	var task []Models.Task
	err := Models.ListTasks(&task)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, task)
	}
}

func GetTask(n *gin.Context) {
	id := n.Params.ByName("id")
	var task Models.Task
	err := Models.GetTask(&task, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, task)
	}
}

func UpdateTask(n *gin.Context) {
	id := n.Params.ByName("id")
	var task Models.Task
	err := Models.GetTask(&task, id)
	if err != nil {
		n.JSON(http.StatusNotFound, task)
	}
	n.BindJSON(&task)
	err = Models.UpdateTask(&task, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, task)
	}
}

func DeleteTask(n *gin.Context) {
	id := n.Params.ByName("id")
	var task Models.Task

	err := Models.DeleteTask(&task, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, id)
	}
}

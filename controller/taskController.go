package controller

import (
	"net/http"

	"github.com/eduardonakaidev/go-tasks/models"
	"github.com/eduardonakaidev/go-tasks/service"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	id := c.Param("idUser")


	var taskRequestDTO models.TaskInputPublic

	if err := c.ShouldBindJSON(&taskRequestDTO); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	err := service.CreateTask(id,taskRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,taskRequestDTO)
}
func Update(c *gin.Context){
	id := c.Param("idUser")


	var taskUpdatedDTO models.TaskInputUpdate
	if err := c.ShouldBindJSON(&taskUpdatedDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	task, err := service.UpdatedTask(id,taskUpdatedDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusOK,task)
}
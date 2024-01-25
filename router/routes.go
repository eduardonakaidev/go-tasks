package router

import (
	"github.com/eduardonakaidev/go-tasks/controller"
	"github.com/gin-gonic/gin"
)

func HandlerRequest() {
	r := gin.Default()
	api := r.Group("/api")
	{
		task := api.Group("/tasks")
		{
			task.POST("/:idUser", controller.CreateTask)
			task.PUT("/:idUser",controller.Update)
		}
	}
}
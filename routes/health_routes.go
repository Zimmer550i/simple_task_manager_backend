package routes

import (
	"github.com/Zimmer550i/simple_task_manager_backend/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine){
	r.GET("/", handlers.HealthCheck)
}
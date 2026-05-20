package main

import (
	"log"

	"github.com/Zimmer550i/simple_task_manager_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterHealthRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Printf("Server dropped: %v", err)
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Initialize() {
	router := gin.Default()
	
	initializeRoutes(router);

	err := router.Run(":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

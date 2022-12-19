package main

import (
	"api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/items", service.GetAllItems)

	router.Run("localhost:8080")
}

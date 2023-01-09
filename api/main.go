package main

import (
	"api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/items", service.GetAllItems)
	router.GET("/items/:item_id", service.GetItemById)

	router.Run("localhost:8080")
}

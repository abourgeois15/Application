package main

import (
	"api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/items", service.GetAllItems)
	router.GET("/items/:item_name", service.GetItemByName)

	router.GET("/machines", service.GetAllMachines)
	router.GET("/machines/:machine_id", service.GetMachineById)

	router.Run("localhost:8080")
}

package main

import (
	"api/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/items", service.GetAllItems)
	router.GET("/items/:item_name", service.GetItemByName)

	router.GET("/machines", service.GetAllMachines)
	router.GET("/machines/:machine_id", service.GetMachineById)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Error when opening server: ", err)
	}
}

package main

import (
	"api/config"
	"api/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectMySQLDB()
	router := gin.Default()
	router.GET("/items", service.GetAllItems)
	router.GET("/item/:item_name", service.GetItemByName)
	router.DELETE("/item/:item_name", service.DeleteItem)
	router.POST("/item", service.CreateItem)
	router.PUT("/item", service.UpdateItem)

	router.GET("/machines", service.GetAllMachines)
	router.GET("/machines/:machine_id", service.GetMachineById)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Origin"}
	router.Use(cors.New(config))

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Error when opening server: ", err)
	}
}

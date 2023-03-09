package router

import (
	"api/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()
	router.POST("/tables", service.CreateTables)

	router.GET("/items", service.GetAllItems)
	router.GET("/item/:item_id", service.GetItemByName)
	router.DELETE("/item/:item_id", service.DeleteItem)
	router.POST("/item", service.CreateItem)
	router.PUT("/item/:item_id", service.UpdateItem)

	router.GET("/machines", service.GetAllMachines)
	router.GET("/machines/type", service.GetAllTypes)
	router.GET("/machine/id/:machine_id", service.GetMachineByID)
	router.GET("/machine/type/:machine_type", service.GetMachineByType)
	router.DELETE("/machine/:machine_id", service.DeleteMachine)
	router.POST("/machine", service.CreateMachine)
	router.PUT("/machine/:machine_id", service.UpdateMachine)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Origin"}
	router.Use(cors.New(config))

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error when opening server: ", err)
	}
}

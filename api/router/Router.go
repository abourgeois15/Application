package router

import (
	"api/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine) {
	router.GET("/items", service.GetAllItems)
	router.POST("/items", service.CreateItem)
	router.GET("/items/:id", service.GetItemByID)
	router.DELETE("/items/:id", service.DeleteItem)
	router.PUT("/items/:id", service.UpdateItem)

	router.GET("/machines", service.GetAllMachines)
	router.GET("/machines?", service.GetMachineByID, service.GetMachineByType)
	router.POST("/machines", service.CreateMachine)
	router.GET("/machines/:id", service.GetMachineByID)
	router.PUT("/machines/:id", service.UpdateMachine)
	router.DELETE("/machines/:id", service.DeleteMachine)

	router.GET("/machine-types", service.GetAllMachineTypes)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Origin"}
	router.Use(cors.New(config))
}

func HandleRequest() {
	router := gin.Default()
	CreateRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error when opening server: ", err)
	}
}

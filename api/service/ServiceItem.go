package service

import (
	"api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config.GetItems())
}

func GetItemByName(c *gin.Context) {
	itemList := config.GetItems()
	name := c.Param("item_name")

	for _, item := range itemList.ItemList {
		if item.Name == name {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func GetAllMachines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config.GetMachines())
}

func GetMachineByName(c *gin.Context) {
	machineList := config.GetMachines()
	name := c.Param("machine_name")

	for _, machine := range machineList.MachineList {
		if machine.Name == name {
			c.IndentedJSON(http.StatusOK, machine)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

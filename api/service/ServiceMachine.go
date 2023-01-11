package service

import (
	"api/config"
	"api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMachines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config.GetMachines())
}

func GetMachineById(c *gin.Context) {
	machineList := config.GetMachines()
	id := c.Param("machine_id")
	var mList []entities.Machine

	for _, machine := range machineList {
		if machine.Name == id || machine.Type == id {
			mList = append(mList, machine)
		}
	}
	if len(mList) > 0 {
		c.IndentedJSON(http.StatusOK, mList)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "machine not found"})
}

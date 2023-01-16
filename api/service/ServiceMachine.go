package service

import (
	"api/config"
	mysqloperations "api/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMachines(c *gin.Context) {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		machineModel := mysqloperations.MachineModel{
			Db: db,
		}
		machines, err := machineModel.FindAll()

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, machines)
	}
}

func GetMachineById(c *gin.Context) {
	db, err := config.GetMySQLDB()
	id := c.Param("machine_id")

	if err != nil {
		fmt.Println(err)
	} else {
		machineModel := mysqloperations.MachineModel{
			Db: db,
		}
		machines, err := machineModel.Find(id)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, machines)
	}
}

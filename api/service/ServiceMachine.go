package service

import (
	"api/config"
	"api/entities"
	mysqloperations "api/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMachine(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdMachine entities.Machine

	if err := c.BindJSON(&createdMachine); err != nil {
		fmt.Println(err)
		return
	}

	machineModel := mysqloperations.MachineModel{
		Db: db,
	}
	machine := entities.Machine{
		Name:   createdMachine.Name,
		Time:   createdMachine.Time,
		Recipe: createdMachine.Recipe,
		Type:   createdMachine.Type,
		Speed:  createdMachine.Speed,
	}
	rows, err := machineModel.Create(&machine)
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusOK, rows)
	}
}

func UpdateMachine(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdMachine entities.Machine

	if err := c.BindJSON(&createdMachine); err != nil {
		return
	}

	machineModel := mysqloperations.MachineModel{
		Db: db,
	}
	machine := entities.Machine{
		Name:   createdMachine.Name,
		Time:   createdMachine.Time,
		Recipe: createdMachine.Recipe,
		Type:   createdMachine.Type,
		Speed:  createdMachine.Speed,
	}
	rows, err := machineModel.Update(machine)
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusOK, rows)
	}
}

func DeleteMachine(c *gin.Context) {
	db, err := config.GetMySQLDB()
	name := c.Param("machine_name")

	if err != nil {
		fmt.Println(err)
	} else {
		machineModel := mysqloperations.MachineModel{
			Db: db,
		}
		rows, err := machineModel.Delete(name)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}
		c.IndentedJSON(http.StatusOK, rows)
	}

}

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

func GetMachineByName(c *gin.Context) {
	db, err := config.GetMySQLDB()
	name := c.Param("machine_name")

	if err != nil {
		fmt.Println(err)
	} else {
		machineModel := mysqloperations.MachineModel{
			Db: db,
		}
		machine, err := machineModel.FindName(name)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, machine)
	}
}

func GetMachineByType(c *gin.Context) {
	db, err := config.GetMySQLDB()
	mtype := c.Param("machine_type")

	if err != nil {
		fmt.Println(err)
	} else {
		machineModel := mysqloperations.MachineModel{
			Db: db,
		}
		machines, err := machineModel.FindType(mtype)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, machines)
	}
}

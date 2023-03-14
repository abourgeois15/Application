package service

import (
	"api/config"
	"api/entities"
	"api/mysql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMachine(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdMachine entities.Machine
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&createdMachine); err != nil {
		fmt.Println(err)
		return
	}

	model := mysql.Model{
		Db: db,
	}
	machine := entities.Machine{
		Name:   createdMachine.Name,
		Time:   createdMachine.Time,
		Recipe: createdMachine.Recipe,
		Type:   createdMachine.Type,
		Speed:  createdMachine.Speed,
	}
	rows, err := model.CreateMachine(&machine)
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusCreated, rows)
	}
}

func UpdateMachine(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdMachine entities.Machine
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&createdMachine); err != nil {
		return
	}

	model := mysql.Model{
		Db: db,
	}
	machine := entities.Machine{
		Id:     createdMachine.Id,
		Name:   createdMachine.Name,
		Time:   createdMachine.Time,
		Recipe: createdMachine.Recipe,
		Type:   createdMachine.Type,
		Speed:  createdMachine.Speed,
	}
	rows, err := model.UpdateMachine(machine)
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
	id, err := strconv.Atoi(c.Param("machine_id"))
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{
			Db: db,
		}
		rows, err := model.DeleteMachine(id)
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
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{
			Db: db,
		}
		machines, err := model.FindAllMachines()

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, machines)
	}
}

func GetAllMachineTypes(c *gin.Context) {
	db, err := config.GetMySQLDB()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{
			Db: db,
		}
		types, err := model.FindAllTypes()

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, types)
	}
}

func GetMachineByID(c *gin.Context) {
	db, err := config.GetMySQLDB()
	id, err := strconv.Atoi(c.Param("id"))
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{
			Db: db,
		}
		machine, err := model.FindMachineById(id)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, machine)
	}
}

func GetMachineByType(c *gin.Context) {
	db, err := config.GetMySQLDB()
	mtype := c.Param("machine_type")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{
			Db: db,
		}
		machines, err := model.FindMachinesByType(mtype)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, machines)
	}
}

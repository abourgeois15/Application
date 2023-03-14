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

func CreateItem(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdItem entities.Item
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&createdItem); err != nil {
		fmt.Println(err)
		return
	}

	model := mysql.Model{Db: db}
	item := entities.Item{
		Name:        createdItem.Name,
		Time:        createdItem.Time,
		Recipe:      createdItem.Recipe,
		Result:      createdItem.Result,
		MachineType: createdItem.MachineType,
	}
	rows, err := model.CreateItem(&item)
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusCreated, rows)
	}
}

func UpdateItem(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdItem entities.Item
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&createdItem); err != nil {
		return
	}

	model := mysql.Model{Db: db}
	item := entities.Item{
		Id:          createdItem.Id,
		Name:        createdItem.Name,
		Time:        createdItem.Time,
		Recipe:      createdItem.Recipe,
		Result:      createdItem.Result,
		MachineType: createdItem.MachineType,
	}
	rows, err := model.UpdateItem(item)
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusOK, rows)
	}
}

func DeleteItem(c *gin.Context) {
	db, err := config.GetMySQLDB()
	id, err := strconv.Atoi(c.Param("id"))
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{Db: db}
		rows, err := model.DeleteItem(id)
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

func GetAllItems(c *gin.Context) {
	db, err := config.GetMySQLDB()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{Db: db}
		items, err := model.FindAllItems()

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, items)
	}
}

func GetItemByID(c *gin.Context) {
	db, err := config.GetMySQLDB()
	id, err := strconv.Atoi(c.Param("id"))
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		model := mysql.Model{Db: db}
		items, err := model.FindItem(id)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, items)
	}
}

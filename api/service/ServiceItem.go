package service

import (
	"api/config"
	"api/entities"
	mysqloperations "api/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTables(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	itemModel := mysqloperations.ItemModel{
		Db: db,
	}

	rows, err := itemModel.CreateTables()
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusCreated, rows)
	}
}

func DeleteTableItem(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	itemModel := mysqloperations.ItemModel{
		Db: db,
	}

	rows, err := itemModel.DeleteTable()
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusOK, rows)
	}
}

func CreateItem(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdItem entities.Item
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&createdItem); err != nil {
		fmt.Println(err)
		return
	}

	itemModel := mysqloperations.ItemModel{
		Db: db,
	}
	item := entities.Item{
		Name:        createdItem.Name,
		Time:        createdItem.Time,
		Recipe:      createdItem.Recipe,
		Result:      createdItem.Result,
		MachineType: createdItem.MachineType,
	}
	rows, err := itemModel.Create(&item)
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

	itemModel := mysqloperations.ItemModel{
		Db: db,
	}
	item := entities.Item{
		Id:          createdItem.Id,
		Name:        createdItem.Name,
		Time:        createdItem.Time,
		Recipe:      createdItem.Recipe,
		Result:      createdItem.Result,
		MachineType: createdItem.MachineType,
	}
	rows, err := itemModel.Update(item)
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
	name := c.Param("item_name")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		itemModel := mysqloperations.ItemModel{
			Db: db,
		}
		rows, err := itemModel.Delete(name)
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
		itemModel := mysqloperations.ItemModel{
			Db: db,
		}
		names, err := itemModel.FindAll()

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, names)
	}
}

func GetItemByName(c *gin.Context) {
	db, err := config.GetMySQLDB()
	name := c.Param("item_name")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		itemModel := mysqloperations.ItemModel{
			Db: db,
		}
		items, err := itemModel.Find(name)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, items)
	}
}

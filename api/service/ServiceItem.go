package service

import (
	"api/config"
	mysqloperations "api/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		itemModel := mysqloperations.ItemModel{
			Db: db,
		}
		items, err := itemModel.FindAll()

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, items)
	}
}

func GetItemByName(c *gin.Context) {
	db, err := config.GetMySQLDB()
	name := c.Param("item_name")

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

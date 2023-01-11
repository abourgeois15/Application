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

	for _, item := range itemList {
		if item.Name == name {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

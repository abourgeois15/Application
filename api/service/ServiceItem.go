package service

import (
	"api/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config.GetData())
}

func GetItemById(c *gin.Context) {
	itemList := config.GetData()
	id := c.Param("item_id")

	for _, item := range itemList.ItemList {
		if strconv.Itoa(item.Id) == id || item.Name == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

package service

import (
	"api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config.GetData())
}

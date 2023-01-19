package main

import (
	"api/config"
	"api/router"
)

func main() {
	config.ConnectMySQLDB()
	router.HandleRequest()
}

package main

import (
	"api/config"
	"api/router"
	"flag"
)

func main() {
	var test = flag.Bool("test", false, "Connect to the test database")
	flag.Parse()
	if *test {
		config.ConnectMySQLDBTest()
	} else {
		config.ConnectMySQLDB()
	}
	router.HandleRequest()
}

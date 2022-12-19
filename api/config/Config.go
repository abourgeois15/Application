package config

import (
	"api/entities"
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetData() entities.ItemList {

	content, err := ioutil.ReadFile("./database.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var data entities.ItemList
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return data
}

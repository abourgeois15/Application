package config

import (
	"api/entities"
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetItems() entities.ItemList {

	content, err := ioutil.ReadFile("./items.json")
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

func GetMachines() entities.MachineList {

	content, err := ioutil.ReadFile("./machines.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var data entities.MachineList
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return data
}

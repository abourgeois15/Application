package config

import (
	"api/entities"
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetItems() []entities.Item {

	content, err := ioutil.ReadFile("./items.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var data entities.ItemList
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return data.ItemList
}

func GetMachines() []entities.Machine {

	content, err := ioutil.ReadFile("./machines.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var data entities.MachineList
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return data.MachineList
}

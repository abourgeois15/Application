package config

import (
	"api/entities"
	"encoding/json"
	"log"
	"os"
)

func GetItems() []entities.Item {

	content, err := os.ReadFile("./items.json")
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

	content, err := os.ReadFile("./machines.json")
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

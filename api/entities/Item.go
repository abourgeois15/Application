package entities

type Ingredient struct {
	Number int    `json:"number"`
	Item   string `json:"item"`
}

type Item struct {
	Name        string        `json:"name"`
	Recipe      [3]Ingredient `json:"recipe"`
	Time        float32       `json:"time"`
	Result      int           `json:"result"`
	MachineType string        `json:"machineType"`
}

type ItemList struct {
	ItemList []Item `json:"itemList"`
}

type Machine struct {
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Recipe [3]Ingredient `json:"recipe"`
	Time   float32       `json:"time"`
	Speed  float32       `json:"speed"`
}

type MachineList struct {
	MachineList []Machine `json:"machineList"`
}

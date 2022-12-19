package entities

type Ingredient struct {
	Number int    `json:"number"`
	Item   string `json:"item"`
}

type Item struct {
	Name        string       `json:"name"`
	Recipe      []Ingredient `json:"recipe"`
	Time        float32      `json:"time"`
	MachineType string       `json:"machineType"`
}

type ItemList struct {
	ItemList []Item `json:"itemList"`
}

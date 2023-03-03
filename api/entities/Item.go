package entities

type Ingredient struct {
	Id     int     `json:"id"`
	Number float32 `json:"number"`
	Item   string  `json:"item"`
}

type Item struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Recipe      []Ingredient `json:"recipe"`
	Time        float32      `json:"time"`
	Result      int          `json:"result"`
	MachineType string       `json:"machineType"`
}

type ItemList struct {
	ItemList []Item `json:"itemList"`
}

type Machine struct {
	Id     int          `json:"id"`
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Recipe []Ingredient `json:"recipe"`
	Time   float32      `json:"time"`
	Speed  float32      `json:"speed"`
}

type MachineList struct {
	MachineList []Machine `json:"machineList"`
}

type CraftPlan struct {
	ParentId      int           `json:"parentId"`
	Item          string        `json:"item"`
	Number        float32       `json:"number"`
	Time          string        `json:"time"`
	Machine       string        `json:"machine"`
	Machines      []string      `json:"machines"`
	NumberMachine float32       `json:"numberMachine"`
	Recipe        [3]Ingredient `json:"recipe"`
}

package entities

type Item struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Recipe      []Ingredient `json:"recipe"`
	Time        float32      `json:"time"`
	Result      int          `json:"result"`
	MachineType string       `json:"machineType"`
}

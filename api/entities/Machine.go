package entities

type Machine struct {
	Id     int          `json:"id"`
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Recipe []Ingredient `json:"recipe"`
	Time   float32      `json:"time"`
	Speed  float32      `json:"speed"`
}
